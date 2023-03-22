package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	markdownlint(ctx, client.Pipeline("markdownlint"))
	yamllint(ctx, client.Pipeline("yamllint"))
}

func markdownlint(ctx context.Context, client *dagger.Client) {
	client.Container().
		From("node:19").
		WithExec([]string{"npm", "install", "-g", "markdownlint-cli@0.33.0"}).
		WithMountedDirectory("/src", client.Host().Directory(".")).
		WithWorkdir("/src").
		WithExec([]string{"markdownlint", "**/*.md"}).
		ExitCode(ctx)
}

func yamllint(ctx context.Context, client *dagger.Client) {
	client.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "add", "--no-cache", "yamllint"}).
		WithMountedDirectory("/src", client.Host().Directory(".")).
		WithWorkdir("/src").
		WithExec([]string{"yamllint", "."}).
		ExitCode(ctx)
}
