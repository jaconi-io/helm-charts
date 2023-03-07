.PHONY: all clean index lint

charts := $(sort $(dir $(wildcard */Chart.yaml)))
packages := $(addprefix out/, $(charts:/=.tgz))

all: $(packages) index

clean:
	rm -rf out/

index: out/index.yaml

lint: $(charts)
	helm lint $(charts)

out/%.tgz: %
	helm package --destination out $<

out/index.yaml: out/
	helm repo index out/
