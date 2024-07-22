.Phony: build-parser
build-parser:
	@go build -o assets/parser cmd/parser/main.go

.Phony: build-skiff
build-skiff:
	@go build -o hookinator cmd/hookinator/main.go

.Phony: build-all
build-all: build-parser build-skiff

.Phony: railcar-build-binaries
railcar-build-binaries:
	goreleaser release --snapshot --clean --config=./assets/.goreleaser.yaml

.Phony: go-test
go-test:
	go test ./... --race
