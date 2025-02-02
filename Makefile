.PHONY: help
help: ## Lists the available commands. Add a comment with '##' to describe a command.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST)\
		| sort\
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the sway-setter cli
	@go build -o bin/sway-setter .

.PHONY: test
test: ## Run all tests
	@go test -v ./...

.PHONY: fmt
fmt: ## Run gofmt
	@gofmt -w .

.PHONY: test-e2e
test-e2e: build ## Run the e2e tests
	@go test ./e2e -v

.PHONY: test-e2e-update
test-e2e-update: build ## Run the e2e tests snapshots and update
	UPDATE_SNAPS=true go test ./e2e -v

.PHONY: nix-flake-check
nix-flake-check: ## Check the nix flake
	@nix flake check

.PHONY: nix-build-source
nix-build-source: ## Build the nix flake
	@nix build .#source

.PHONY: nix-build-default
nix-build-default: ## Build the nix flake
	@nix build .#

.PHONY: git-hook-precommit
git-hook-precommit: ## Install the pre-commit git hook
	./ci/pre-commit.sh
