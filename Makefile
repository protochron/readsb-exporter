REPO := protochron/readsb-exporter
GIT_SHA := $(shell git rev-parse --short HEAD)
GOARCH ?= amd64

ARCHS := amd64 arm
.PHONY: build
build:
	go build ./cmd/readsb-exporter

release-docker:
	docker buildx build --platform linux/arm/v7,linux/amd64 -t protochron/readsb-exporter:latest --push .

.PHONY: clean
clean:
	@rm -f readsb-exporter
