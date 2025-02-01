.PHONY: help
help: ## Lists the available commands. Add a comment with '##' to describe a command.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST)\
		| sort\
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the sway-setter cli
	@go build -o bin/sway-setter cmd/sway-setter/*.go

.PHONY: run
run: ## Run the sway-setter cli
	@go run cmd/sway-setter/*.go

.PHONY: test
test: ## Run all tests
	@go test -v ./...

.PHONY: fmt
fmt: ## Run gofmt
	@gofmt -w .
