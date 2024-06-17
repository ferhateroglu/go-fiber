.PHONY: run test

run:
	go run ./cmd/api/main.go

test:
	gotestsum