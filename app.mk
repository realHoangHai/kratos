GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO              := GO111MODULE=on go

ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"

ifeq ($(OS),Windows_NT)
    IS_WINDOWS:=1
endif

APP_VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
APP_DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "restroom-system/" $$0 ":0.1.0"}')


.PHONY: dep vendor build clean docker gen ent wire buf openapi run test cover vet lint app

# download dependencies of module
dep:
	@go mod download

# create vendor
vendor:
	@go mod vendor

# build golang application
build:
ifeq ("$(wildcard ./bin/)","")
	mkdir bin
endif
	@go build -ldflags "-X main.version=$(APP_VERSION)" -o ./bin/ ./...

# clean build files
clean:
	@go clean
	$(if $(IS_WINDOWS), del "coverage.out", rm -f "coverage.out")

# build docker image
docker:
	@docker build -t $(APP_DOCKER_IMAGE) . \
				  -f ../../.docker/Dockerfile \
				  --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) GRPC_PORT=9000 REST_PORT=8000

# generate code
gen: ent wire buf openapi

# generate ent code
ent:
ifneq ("$(wildcard ./internal/data/ent)","")
	@go run -mod=mod entgo.io/ent/cmd/ent generate \
				--feature privacy \
				--feature sql/modifier \
				--feature entql \
				--feature sql/upsert \
				./internal/data/ent/schema
endif

# generate wire code
wire:
	@go run -mod=mod github.com/google/wire/cmd/wire ./cmd/server

# generate protobuf api go code
buf:
	@cd ../../api && \
	buf generate

# generate OpenAPI v3 doc
openapi:
	@cd ../../api && \
	buf generate --path proto/admin/v1 --template proto/admin/v1/openapi.gen.yaml && \
	buf generate --path proto/client/v1 --template proto/client/v1/openapi.gen.yaml

# run application
run: buf openapi
	@go run ./cmd/server -conf ./configs

# run tests
test:
	@go test ./...

# run coverage tests
cover:
	@go test -v ./... -coverprofile=coverage.out

# run static analysis
vet:
	@go vet

# run lint
lint:
	@golangci-lint run

# build service app
app: buf wire ent build

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help