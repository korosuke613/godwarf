.PHONY: build

build: build/godwarf

build/godwarf: $(wildcard *.go)
	go build -o $@
