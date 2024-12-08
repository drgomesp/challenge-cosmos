BINARY_NAME := rpsd
BUILD_DIR := ./build

# don't override user values
VERSION ?= "coderbyte"

# Update the ldflags with the app, client & server names
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=rps \
	-X github.com/cosmos/cosmos-sdk/version.AppName=rpsd \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

###########
# Install #
###########

all: install

install:
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
	@echo "--> Installing rpsd"
	@go install $(BUILD_FLAGS) -mod=readonly ./cmd/rpsd

init:
	@./scripts/init.sh || { echo "Init script failed"; exit 1; }

##################
###  Protobuf  ###
##################

protoVer=0.15.1
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
DOCKER ?= docker
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating protobuf files..."
	@$(protoImage) sh ./scripts/protocgen.sh
	@go mod tidy

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint proto/ --error-format=json

.PHONY: proto-all proto-gen proto-format proto-lint

#################
###  Linting  ###
#################

golangci_lint_cmd=golangci-lint
golangci_version=v1.51.2

lint: golangci-install
	@echo "--> Running linter"
	@$(golangci_lint_cmd) run ./... --timeout 15m

lint-fix: golangci-install
	@echo "--> Running linter and fixing issues"
	@$(golangci_lint_cmd) run ./... --fix --timeout 15m

golangci-install:
	@echo "--> Installing golangci-lint if not present"
	@command -v $(golangci_lint_cmd) >/dev/null 2>&1 || go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)

.PHONY: lint lint-fix golangci-install
