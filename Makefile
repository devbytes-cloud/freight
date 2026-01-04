.Phony: build-freight
build-freight:
	@go build -o freight cmd/freight/main.go

.Phony: build-all
build-all: conductor-build-binaries build-freight

.Phony: conductor-build-binaries
conductor-build-binaries:
	goreleaser release --snapshot --clean --config=./assets/.goreleaser.yaml

.Phony: conductor-dev-build-binaries
conductor-dev-build-binaries:
	goreleaser release --snapshot --clean --config=./assets/.goreleaser.yaml

.Phony: test-go
test-go:
	go test ./... --race

.Phony: lint
lint:
	golangci-lint run


.Phony: asset-placeholders
asset-placeholders:
	mkdir -p assets/dist/conductor_darwin_amd64_v1 \
		assets/dist/conductor_darwin_arm64_v8.0 \
		assets/dist/conductor_linux_amd64_v1 \
		assets/dist/conductor_linux_arm64_v8.0 \
		assets/dist/conductor_linux_arm_6 \
		assets/dist/conductor_windows_amd64_v1 \
		assets/dist/conductor_windows_arm64_v8.0 \
		assets/dist/conductor_windows_arm_6
	touch assets/dist/conductor_darwin_amd64_v1/conductor \
		assets/dist/conductor_darwin_arm64_v8.0/conductor \
		assets/dist/conductor_linux_amd64_v1/conductor \
		assets/dist/conductor_linux_arm64_v8.0/conductor \
		assets/dist/conductor_linux_arm_6/conductor \
		assets/dist/conductor_windows_amd64_v1/conductor.exe \
		assets/dist/conductor_windows_arm64_v8.0/conductor.exe \
		assets/dist/conductor_windows_arm_6/conductor.exe
