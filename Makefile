.Phony: build-parser
build-parser:
	@go build -o assets/parser cmd/parser/main.go

.Phony: build-skiff
build-skiff:
	@go build -o hookinator cmd/hookinator/main.go

.Phony: build-all
build-all: build-parser build-skiff

.Phony: parser-build-binaries
parser-build-binaries:
	goreleaser release --snapshot --clean --config=./assets/.goreleaser.yaml