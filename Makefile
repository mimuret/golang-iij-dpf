PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
types_go_files  = $(wildcard ./pkg/types/*.go)
v1_core_go_files  = $(wildcard ./pkg/apis/dpf/v1/core/*.go)
v1_contracts_go_files  = $(wildcard ./pkg/apis/dpf/v1/contracts/*.go)
v1_common_configs_go_files  = $(wildcard ./pkg/apis/dpf/v1/common_configs/*.go)
v1_zones_go_files  = $(wildcard ./pkg/apis/dpf/v1/zones/*.go)
v1_lb_domains_go_files  = $(wildcard ./pkg/apis/dpf/v1/lb_domains/*.go)
deep_copy_files = pkg/types/ZZ_deepcopy_generated.go \
	pkg/apis/dpf/v1/core/ZZ_deepcopy_generated.go \
	pkg/apis/dpf/v1/contracts/ZZ_deepcopy_generated.go \
	pkg/apis/dpf/v1/zones/ZZ_deepcopy_generated.go \
	pkg/apis/dpf/v1/common_configs/ZZ_deepcopy_generated.go \
	pkg/apis/dpf/v1/lb_domains/ZZ_deepcopy_generated.go

all: deepcopy fmt checks test

checks: golangci-lint
	$(GOLANGCI_LINT) run

test:
	go test -coverprofile="cover.out" ./...
	go tool cover -html=cover.out -o cover.html

fmt: gofumpt
	$(GOFUMPT) -l -w .

deepcopy: $(deep_copy_files)

pkg/types/ZZ_deepcopy_generated.go: deepcopy-gen $(types_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/types
pkg/apis/dpf/v1/core/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_core_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/dpf/v1/core
pkg/apis/dpf/v1/contracts/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_contracts_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/dpf/v1/contracts
pkg/apis/dpf/v1/zones/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_common_configs_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/dpf/v1/zones
pkg/apis/dpf/v1/common_configs/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_zones_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/dpf/v1/common_configs
pkg/apis/dpf/v1/lb_domains/ZZ_deepcopy_generated.go: deepcopy-gen $(v1_lb_domains_go_files)
	$(DEEPCOPY_GEN) -O ZZ_deepcopy_generated -h boilerplate.go.txt -i ./pkg/apis/dpf/v1/lb_domains

DEEPCOPY_GEN = $(PROJECT_DIR)/bin/deepcopy-gen
deepcopy-gen: ## Download deepcopy-gen locally if necessary.
	mkdir -p $(PROJECT_DIR)/bin
	$(call go-get-tool,$(DEEPCOPY_GEN),k8s.io/code-generator/cmd/deepcopy-gen@latest)

GOFUMPT = $(PROJECT_DIR)/bin/gofumpt
gofumpt:
	mkdir -p $(PROJECT_DIR)/bin
	$(call go-get-tool,$(GOFUMPT),mvdan.cc/gofumpt@latest)

GOLANGCI_LINT = $(PROJECT_DIR)/bin/golangci-lint
golangci-lint:
	mkdir -p $(PROJECT_DIR)/bin
	$(call go-get-tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/cmd/golangci-lint@master)


# go-get-tool will 'go get' any package $2 and install it to $1.
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
}
endef
