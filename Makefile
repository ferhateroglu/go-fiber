.PHONY: run test coverage

run:
	go run ./cmd/api/main.go

test:
	gotestsum

coverage:
	@echo "Running tests with coverage..."
	@gotestsum -- -coverprofile=coverage.out ./...
	@echo "\nCoverage summary:"
	@go tool cover -func=coverage.out | grep total:
	@echo "\nGenerating HTML report..."
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@xdg-open coverage.html 