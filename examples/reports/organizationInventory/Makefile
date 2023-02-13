BINARY=organizationReport
VERSION=0.0.1

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(BINARY)-linux-amd64 -ldflags "-X main.version=$(VERSION)"
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/$(BINARY)-linux-arm64 -ldflags "-X main.version=$(VERSION)"

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/$(BINARY)-windows-amd64.exe -ldflags "-X main.version=$(VERSION)"
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/$(BINARY)-windows-arm64.exe -ldflags "-X main.version=$(VERSION)"

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/$(BINARY)-macos-amd64 -ldflags "-X main.version=$(VERSION)"
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/$(BINARY)-macos-arm64 -ldflags "-X main.version=$(VERSION)"

.PHONY: clean
clean:
	rm -rf build/*