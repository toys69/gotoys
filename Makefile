lint:
	go version
	go mod tidy
	gofumpt -w .
	goimports -l -w .
	gofmt -l -s -w .
	go vet ./...
	golangci-lint run --fix ./...
