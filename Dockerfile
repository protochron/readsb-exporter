FROM --platform=${BUILDPLATFORM} golang:1.15.2 as build
WORKDIR /readsb-exporter
ENV CGO_ENABLED=0
COPY . /readsb-exporter
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags "-s -w" ./cmd/readsb-exporter

FROM --platform=${BUILDPLATFORM} scratch
COPY --from=build /readsb-exporter/readsb-exporter /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/readsb-exporter"]
