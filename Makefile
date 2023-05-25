export

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

test: ### run test localy
	go test -v -cover -race .
.PHONY: test

bench: ## run benchmarks localy
	go test -benchmem -bench .
.PHONY: bench

cbench: ## run benchmarks in container
	docker build . -t dist && docker run --rm dict
.PHONY: cbench