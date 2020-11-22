.PHONY: build test test-integration

build:
	@echo "Building binary unit"
	@go build -o slowly main.go

test:
	@echo "Testing unit"
	@go test ./app/... -cover -count=1

test-integration:
	@echo "Testing integration"
	@go test ./test/... -count=1
