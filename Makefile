.DEFAULT_GOAL = help

TOOLS_DIR := ./tools
TOOLS_BIN := $(TOOLS_DIR)/bin

export PATH := $(TOOLS_BIN):$(PATH)

GO_TOOLS := github.com/bufbuild/buf/cmd/buf \
            github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
            github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
            github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
            google.golang.org/grpc/cmd/protoc-gen-go-grpc \
            google.golang.org/protobuf/cmd/protoc-gen-go

GO_TOOLS_BIN := $(addprefix $(TOOLS_BIN), $(notdir $(GO_TOOLS)))
GO_TOOLS_VENDOR := $(addprefix vendor/, $(GO_TOOLS))

CONTENT_V2_PROTO  := proto/content/v2
CONTENT_V2_GEN_GO := golang/content/v2

$(TOOLS_DIR):
	mkdir -v -p $@

$(GO_TOOLS_BIN): $(GO_TOOLS_VENDOR)
	GOBIN="$(PWD)/$(TOOLS_BIN)" go install -mod=vendor $(GO_TOOLS)

## help: List available make targets and descriptions
.PHONY: help
help: Makefile
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

vendor:
	go mod vendor

## tools-install: Reinstall build tools
.PHONY: tools-install
tools-install: vendor $(TOOLS_DIR) $(GO_TOOLS_BIN)

## tools-upgrade: Upgrade build tools
.PHONY: tools-upgrade
tools-upgrade:
	@rm -rf $(PWD)/vendor
	@echo $(GO_TOOLS) | xargs -n1 go get -u $$1
	@go mod vendor

BUF := $(TOOLS_BIN)/buf

.PHONY: content-v2-service-article
content-v2-service-article:
	@$(BUF) generate --path $(CONTENT_V2_PROTO)/article.proto
	@mv $(CONTENT_V2_GEN_GO)/article* $(CONTENT_V2_GEN_GO)/

## generate: Generate stubs for protoc plugins
.PHONY: generate
generate: content-v2-service-article

## breaking: Check if schema has breaking changes
.PHONY: breaking
breaking:
	@$(BUF) breaking --against '.git#branch=main'
