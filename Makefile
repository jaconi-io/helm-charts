.PHONY: all clean index lint

charts := $(sort $(dir $(wildcard */Chart.yaml)))
packages := $(addprefix out/, $(charts:/=.tgz))

all: $(packages) index

clean:
	rm -rf out/

index: out/index.html out/index.yaml out/artifacthub-repo.yml

lint: $(charts)
	helm lint $(charts)

out/%.tgz: %
	helm package --destination out $<

out/index.yaml: out/
	helm repo index out/

out/index.html: index.html
	cp index.html out/

out/artifacthub-repo.yml: artifacthub-repo.yml
	cp artifacthub-repo.yml out/
