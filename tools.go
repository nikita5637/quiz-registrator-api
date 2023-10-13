//go:build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/xo/xo"
	_ "golang.org/x/tools/cmd/goimports"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "gotest.tools/gotestsum"
)
