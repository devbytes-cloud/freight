.Phony: build-freight
build-freight:
	@go build -o freight cmd/freight/main.go

.Phony: build-all
build-all: railcar-build-binaries build-freight

.Phony: railcar-build-binaries
railcar-build-binaries:
	goreleaser release --snapshot --clean --config=./assets/.goreleaser.yaml

.Phony: test-go
test-go:
	go test ./... --race

.Phony: lint
lint:
	golangci-lint run