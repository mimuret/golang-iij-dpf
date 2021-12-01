types_go_files  = $(wildcard ./pkg/types/*.go)
v1_core_go_files  = $(wildcard ./pkg/apis/dpf/v1/core/*.go)
v1_contracts_go_files  = $(wildcard ./pkg/apis/dpf/v1/contracts/*.go)
v1_common_configs_go_files  = $(wildcard ./pkg/apis/dpf/v1/common_configs/*.go)
v1_zones_go_files  = $(wildcard ./pkg/apis/dpf/v1/zones/*.go)
deep_copy_files = pkg/types/ZZ_deepcopy_generated.go \
	pkg/apis/dpf/v1/core/ZZ_deepcopy_generated.go \
	pkg/apis/dpf/v1/contracts/ZZ_deepcopy_generated.go \
	pkg/apis/dpf/v1/zones/ZZ_deepcopy_generated.go \
	pkg/apis/dpf/v1/common_configs/ZZ_deepcopy_generated.go

all: deepcopy checks test fmt

checks:
	golangci-lint run

test:
	go test -coverprofile="cover.out" ./...
	go tool cover -html=cover.out -o cover.html

fmt:
	go fmt ./...

deepcopy:  

pkg/types/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_core_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/types
pkg/apis/dpf/v1/core/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_core_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/dpf/v1/core
pkg/apis/dpf/v1/contracts/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_contracts_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/dpf/v1/contracts
pkg/apis/dpf/v1/zones/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_common_configs_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/dpf/v1/zones
pkg/apis/dpf/v1/common_configs/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_zones_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/dpf/v1/common_configs

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
