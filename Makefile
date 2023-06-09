LOCAL_BIN:=$(CURDIR)/bin
GOIMPORTS_BIN:=$(LOCAL_BIN)/goimports
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
PROTOC_GEN_GO_BIN:=$(LOCAL_BIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC_BIN:=$(LOCAL_BIN)/protoc-gen-go-grpc
XO_BIN:=$(LOCAL_BIN)/xo

.PHONY: .install-bin-deps
.install-bin-deps:
ifeq ($(wildcard $(PROTOC_GEN_GO_BIN)),)
	$(info Installing binary dependency protoc-gen-go)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go
endif
ifeq ($(wildcard $(PROTOC_GEN_GO_GRPC_BIN)),)
	$(info Installing binary dependency protoc-gen-go-grpc)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc 
endif
ifeq ($(wildcard $(GOIMPORTS_BIN)),)
	$(info Installing binary dependency goimports)
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports 
endif

.PHONY: .install-lint
.install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint)
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint
	GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: .install-xo
.install-xo:
ifeq ($(wildcard $(XO_BIN)),)
	$(info Downloading xo)
	GOBIN=$(LOCAL_BIN) go install github.com/xo/xo
	XO_BIN:=$(LOCAL_BIN)/xo
endif

.PHONY: build
build:
	go build -o registrator-api ./cmd/registrator-api

.PHONY: clean
clean:
	rm registrator-api

.PHONY: gen-db
gen-db: .install-xo
	$(XO_BIN) schema \
	--schema="" \
	--go-import="github.com/nikita5637/quiz-registrator-api/internal/pkg/logger" \
	--go-import="github.com/nikita5637/quiz-registrator-api/internal/pkg/tx" \
	--out internal/pkg/storage/mysql \
	--src templates \
	--template go \
	mysql://$(USERNAME):$(DATABASE_PASSWORD)@$(DBADDR)/$(DBNAME)
	rm -rf ./internal/pkg/storage/mysql/goosedbversion.xo.go ./internal/pkg/storage/mysql/db.xo.go

.PHONY: generate
generate: .install-bin-deps
	PATH="$(LOCAL_BIN):$(PATH)" && protoc \
	--go_out=./pkg/pb --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/pb --go-grpc_opt=paths=source_relative \
	--proto_path=./api ./api/certificate_manager/certificate_manager.proto

	PATH="$(LOCAL_BIN):$(PATH)" && protoc \
	--go_out=./pkg/pb --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/pb --go-grpc_opt=paths=source_relative \
	--proto_path=./api ./api/registrator/registrator.proto

	$(GOIMPORTS_BIN) -w ./

.PHONY: go-generate
go-generate:
	go generate ./...

.PHONY: lint
lint: .install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --config=.golangci.pipeline.yaml ./...

.PHONY: migrations
migrations:
	goose -dir migrations/ mysql "$(USERNAME):$(DATABASE_PASSWORD)@tcp($(DBADDR))/$(DBNAME)" up

.PHONY: run
run:
	go run ./cmd/registrator-api

.PHONY: test
test:
	go test -v ./...
