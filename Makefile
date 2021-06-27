REPO := ghcr.io/protochron/readsb-exporter
GIT_SHA := $(shell git rev-parse --short HEAD)
GOARCH ?= amd64

.PHONY: build
build:
	go build ./cmd/readsb-exporter

.PHONY: release-docker
release-docker:
	docker buildx build --platform linux/arm/v7,linux/amd64,linux/arm64 -t $(REPO):$(GIT_SHA) --push .

.PHONY: clean
clean:
	@rm -f readsb-exporter
