GO ?= go

.PHONY: proto install build profile hedone

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

init: ## Initialise new module project in current directory. Use "mod=" flag to specify module name. e.g. mod=github.com/linushung/hedone
	${GO} mod init ${mod}

tidy: ## Add missing and remove unused modules
	${GO} mod tidy -v

build: ## Build go program
	${GO} build

aletheia: build ## Run aletheia program
	./aletheia
