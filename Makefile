.PHONY: build release

GEN_PATH := release/gen/protobuf/go

build:
	@cmake --preset=build
	@cmake --build build

release:
	@cmake --preset=release
	@cmake --build release

gen-go:
	@rm -rf "$(GEN_PATH)"
	@cmake --preset=release -DGEN_GO=ON
	@cp -r $(GEN_PATH)/*  go/

mod:
	@go mod tidy
	@go mod download
