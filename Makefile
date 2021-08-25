LOCAL_BIN:=$(CURDIR)/bin

.PHONY: deps
deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
	ls go.mod || go mod init gitlab.com/siriusfreak/ova-entertainment-api
	GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: generate
generate:
	GOBIN=$(LOCAL_BIN) protoc -I api/ \
	  --go_out=pkg/ova-entertainment-api --go_opt=paths=import \
	  --go-grpc_out=pkg/ova-entertainment-api --go-grpc_opt=paths=import \
	  api/ova-entertainment-api/*.proto

.PHONY: test
test:
	GOBIN=$(LOCAL_BIN) go test ./...

.PHONY: run
run:
	GOBIN=$(LOCAL_BIN) go run cmd/ova-entertainment-api/main