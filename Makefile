.Phony: build-freight
build-freight:
	@go build -o freight cmd/freight/main.go

.Phony: build-all
build-all: conductor-build-binaries build-freight

.Phony: flox-build
flox-build:
	flox build freight

CONDUCTOR_BUILD := CGO_ENABLED=0 go build -trimpath -ldflags "-s -w"

.Phony: conductor-build-binaries
conductor-build-binaries:
	rm -rf assets/dist
	GOOS=darwin  GOARCH=amd64         $(CONDUCTOR_BUILD) -o assets/dist/conductor_darwin_amd64_v1/conductor        ./cmd/conductor
	GOOS=darwin  GOARCH=arm64         $(CONDUCTOR_BUILD) -o assets/dist/conductor_darwin_arm64_v8.0/conductor      ./cmd/conductor
	GOOS=linux   GOARCH=amd64         $(CONDUCTOR_BUILD) -o assets/dist/conductor_linux_amd64_v1/conductor         ./cmd/conductor
	GOOS=linux   GOARCH=arm64         $(CONDUCTOR_BUILD) -o assets/dist/conductor_linux_arm64_v8.0/conductor       ./cmd/conductor
	GOOS=linux   GOARCH=arm   GOARM=6 $(CONDUCTOR_BUILD) -o assets/dist/conductor_linux_arm_6/conductor            ./cmd/conductor
	GOOS=windows GOARCH=amd64         $(CONDUCTOR_BUILD) -o assets/dist/conductor_windows_amd64_v1/conductor.exe   ./cmd/conductor
	GOOS=windows GOARCH=arm64         $(CONDUCTOR_BUILD) -o assets/dist/conductor_windows_arm64_v8.0/conductor.exe ./cmd/conductor

.Phony: conductor-dev-build-binaries
conductor-dev-build-binaries: conductor-build-binaries

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
		assets/dist/conductor_windows_arm64_v8.0
	touch assets/dist/conductor_darwin_amd64_v1/conductor \
		assets/dist/conductor_darwin_arm64_v8.0/conductor \
		assets/dist/conductor_linux_amd64_v1/conductor \
		assets/dist/conductor_linux_arm64_v8.0/conductor \
		assets/dist/conductor_linux_arm_6/conductor \
		assets/dist/conductor_windows_amd64_v1/conductor.exe \
		assets/dist/conductor_windows_arm64_v8.0/conductor.exe
