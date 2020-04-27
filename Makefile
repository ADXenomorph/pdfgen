.PHONY: build test help

build: ## Build the binary
	go build cmd/pdfgen/pdfgen.go

test: ## Run tests
	go clean -testcache && go test ./...

install: build ## Install as a binary
	go install cmd/pdfgen/pdfgen.go

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

