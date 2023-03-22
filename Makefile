.PHONY: all clean index lint

MAKEFLAGS += rR

charts := $(sort $(dir $(wildcard */Chart.yaml)))
packages := $(addprefix out/, $(charts:/=.tgz))

all: $(packages) index

clean:
	rm -rf out/

index: out/CNAME out/index.html out/index.yaml out/artifacthub-repo.yml

lint: $(charts)
	helm lint $(charts)

out/%.tgz: %
	helm package --destination out $<

out/index.yaml: $(packages)
	helm repo index out/

out/%: %
	mkdir -p $(dir $@)
	cp $< $@

# Disable builtin rules.
.SUFFIXES:
