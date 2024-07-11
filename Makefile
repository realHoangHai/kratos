GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
CURRENT_DIR:=$(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
ROOT_DIR:=$(dir $(realpath $(lastword $(MAKEFILE_LIST))))
SRCS_MK:=$(foreach dir, app, $(wildcard $(dir)/*/*/Makefile))

ifeq ($(OS),Windows_NT)
    IS_WINDOWS:=1
endif

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
	API_GO_FILES=$(shell $(Git_Bash) -c "find api -name *.go")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
	API_GO_FILES=$(shell find api -name *.go)
endif








.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/gnostic@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/google/wire/cmd/wire@latest
	go install entgo.io/ent/cmd/ent@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
 	       --go-errors_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

# adding proto files
kratos.add:
	kratos proto add api/helloworld/demo.proto

# generate proto codes
kratos.client:
	kratos proto client api/helloworld/demo.proto

# generate service codes
kratos.server:
	kratos proto server api/helloworld/demo.proto -t internal/service

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: buf buf.update
# buf
buf:
	rm -rf api/gen/ && \
	cd api && \
	buf generate

buf.update:
	cd api && \
	buf dep update

buf.lint:
	cd api && \
	buf lint

.PHONY: ent
# ent
ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier ./internal/data/ent/schema

.PHONY: wire
# wire
wire:
	cd cmd/kratos && \
	wire

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	rm -rf $(API_GO_FILES)
	make api;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
