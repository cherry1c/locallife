PROJECT_NAME := "local-life"
PKG := "."
PKG_LIST := $(shell go mod tidy && go list ${PKG}/...)
BUS_PKG_LIST := $(shell go mod tidy && go list ${PKG}/... | grep -v unittest)
GO_FILES := $(shell find . -name '*.go' | grep -v _test.go)

.DEFAULT_GOAL := default
.PHONY: all

dep: ## Get dependencies
	@echo "go dep..."
	@go mod tidy

fmt: dep ## Format code
	@echo "go fmt..."
	@go fmt $(PKG_LIST)

test: dep ## Run unittests
	@echo "go test..."
	@go test -gcflags=all=-l -short -v -count=1 ${BUS_PKG_LIST}

build: dep fmt ## Build frpc project
	@echo "go build..."
	@CGO_ENABLED=1 go build -v -buildmode=default -o bin/${PROJECT_NAME} main.go
#	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -gcflags=all="-N -l" -o bin/${PROJECT_NAME} cmd/main.go
	@chmod +x bin/${PROJECT_NAME}

run: dep fmt ## Run frpc project
	@echo "go run..."
	@go run main.go



