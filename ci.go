package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
)

const out = "out"

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	markdownlint(ctx, client.Pipeline("markdownlint"))
	yamllint(ctx, client.Pipeline("yamllint"))

	helm := client.Pipeline("helm")
	helmLint(ctx, helm)
	helmPackage(ctx, helm)
	helmRepoIndex(ctx, helm)

	// Copy supporting files.
	client.Host().File(".nojekyll").Export(ctx, filepath.Join(out, ".nojekyll"))
	client.Host().File("artifacthub-repo.yml").Export(ctx, filepath.Join(out, "artifacthub-repo.yml"))
	client.Host().File("CNAME").Export(ctx, filepath.Join(out, "CNAME"))
	client.Host().File("index.html").Export(ctx, filepath.Join(out, "index.html"))
}

func markdownlint(ctx context.Context, client *dagger.Client) {
	exit, err := client.Container().
		From("node:20").
		WithExec([]string{"npm", "install", "-g", "markdownlint-cli@0.35.0"}).
		WithMountedDirectory("/src", client.Host().Directory(".")).
		WithWorkdir("/src").
		WithExec([]string{"markdownlint", "**/*.md"}).
		ExitCode(ctx)
	if err != nil {
		panic(fmt.Errorf("[markdownlint] %w", err))
	}

	if exit != 0 {
		panic(fmt.Errorf("[markdownlint] unexpected exit code: %d", exit))
	}
}

func yamllint(ctx context.Context, client *dagger.Client) {
	exit, err := client.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "add", "--no-cache", "yamllint"}).
		WithMountedDirectory("/src", client.Host().Directory(".")).
		WithWorkdir("/src").
		WithExec([]string{"yamllint", "."}).
		ExitCode(ctx)
	if err != nil {
		panic(fmt.Errorf("[yamllint] %w", err))
	}

	if exit != 0 {
		panic(fmt.Errorf("[yamllint] unexpected exit code: %d", exit))
	}
}

func helmLint(ctx context.Context, client *dagger.Client) {
	exit, err := client.Container().
		From("alpine/helm").
		WithMountedDirectory("/src", client.Host().Directory(".")).
		WithWorkdir("/src").
		WithExec(append([]string{"lint"}, charts(ctx, client)...)).
		ExitCode(ctx)
	if err != nil {
		panic(fmt.Errorf("[helm lint] %w", err))
	}

	if exit != 0 {
		panic(fmt.Errorf("[helm lint] unexpected exit code: %d", exit))
	}
}

func helmPackage(ctx context.Context, client *dagger.Client) {
	tmp := client.Directory()

	for _, chart := range charts(ctx, client) {
		pkg := client.Container().
			From("alpine/helm").
			WithMountedDirectory("/src", client.Host().Directory(chart)).
			WithWorkdir("/src").
			WithExec([]string{"package", "--destination", "/out", "."})

		tmp = tmp.WithDirectory(".", pkg.Directory("/out"))
	}

	ok, err := tmp.Export(ctx, out)
	if err != nil {
		panic(fmt.Errorf("[helm package] %w", err))
	}

	if !ok {
		panic(errors.New("[helm package] package not found"))
	}
}

func helmRepoIndex(ctx context.Context, client *dagger.Client) {
	idx := client.Container().
		From("alpine/helm").
		WithMountedDirectory("/src", client.Host().Directory(out)).
		WithWorkdir("/src").
		WithExec([]string{"repo", "index", "."}).
		File("index.yaml")

	ok, err := idx.Export(ctx, filepath.Join(out, "index.yaml"))
	if err != nil {
		panic(fmt.Errorf("[helm repo index] %w", err))
	}

	if !ok {
		panic(errors.New("[helm repo index] index.yaml not found"))
	}
}

// charts returns a slice of all Helm chart directories (i.e. directories containing a Chart.yaml file).
func charts(ctx context.Context, client *dagger.Client) []string {
	charts, err := client.Host().Directory(".", dagger.HostDirectoryOpts{
		Include: []string{"*/Chart.yaml"},
	}).Entries(ctx)
	if err != nil {
		panic(err)
	}
	return charts
}
