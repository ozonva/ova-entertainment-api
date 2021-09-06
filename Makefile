include .env

LOCAL_BIN:=$(CURDIR)/bin
DBSTRING:="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5434/$(POSTGRES_DB)?sslmode=disable"

.PHONY: build
build: deps
	GOBIN=$(LOCAL_BIN) go build -o $(LOCAL_BIN)/ova-entertainment-api cmd/ova-entertainment-api/main.go

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
	GOBIN=$(LOCAL_BIN) go test -v ./...

.PHONY: run
run:
	GOBIN=$(LOCAL_BIN) go run cmd/ova-entertainment-api/main

.PHONY: migrate-up
migrate-up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir migration status
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir migration up

#for fun @todo manualcov
.PHONY: create-badge
create-badge:
	GOBIN=$(LOCAL_BIN) go test ./... -coverprofile cover.out.tmp
	GOBIN=$(LOCAL_BIN) go tool cover -html=cover.out.tmp -o cover.html
	cat cover.out.tmp | grep -v "mock_" > coverage.out
	GOBIN=$(LOCAL_BIN) go tool cover -func coverage.out | awk 'END {print $3+0}'
	rm cover.out.tmp
	rm coverage.out
	#gopherbadger -md="README.md" -manualcov=

.PHONY: docs
docs:
	godoc -http=:6060

.PHONY: gosec
gosec:
	gosec -exclude=G102 ./...

.PHONY: mockgen
mockgen:
	cd internal/repo && mockgen -source=repo.go  -destination=mock_repo.go -package=repo
	cd internal/metrics && mockgen -source=metrics.go  -destination=mock_metrics.go -package=metrics
	cd internal/kafka && mockgen -source=kafka.go  -destination=mock_kafka.go -package=kafka

.PHONY: lint
lint:
	gosec -exclude=G102 ./...
	go vet ./...
	go fmt ./...
	golangci-lint run ./...