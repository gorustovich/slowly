.PHONY: test test-integration

test:
	@echo "Testing unit"
	@go test ./app/... -cover -count=1

test-integration:
	@echo "Testing integration"
	@go test ./test/... -count=1
