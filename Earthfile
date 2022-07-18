# Earthfile

VERSION --use-host-command 0.6
FROM golang:1.18
WORKDIR /work

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    FROM +deps
    COPY *.go Makefile .
    RUN make build
    SAVE ARTIFACT build/godwarf /godwarf AS LOCAL build/godwarf

docker:
    FROM ubuntu:22.04
    RUN apt-get update \
     && apt-get install -y --no-install-recommends \
        git \
        ca-certificates \
        golang-go \
     && apt-get -y clean \
     && rm -rf /var/lib/apt/lists/*
    COPY +build/godwarf /
    ENTRYPOINT ["/godwarf"]
    SAVE IMAGE korosuke613/godwarf

demo:
    FROM +docker
    WORKDIR /work
    RUN git clone https://github.com/korosuke613/playground
    COPY demo.yaml .
    ENTRYPOINT ["/godwarf", "./demo.yaml"]
    SAVE IMAGE korosuke613/godwarf:demo
