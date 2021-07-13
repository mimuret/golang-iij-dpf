all: deepcopy test fmt

test:
	go test -race -coverprofile="coverage.txt" -covermode=atomic ./...

fmt:
	go fmt ./...

deepcopy: deepcopy-gen
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/api
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/types
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/core
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/contracts
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/zones
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/common_configs

DEEPCOPY_GEN = $(shell pwd)/bin/deepcopy-gen
deepcopy-gen: ## Download deepcopy-gen locally if necessary.
	mkdir -p $(shell pwd)/bin
	$(call go-get-tool,$(DEEPCOPY_GEN),k8s.io/code-generator/cmd/deepcopy-gen)

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
