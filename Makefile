PROTO_DIR := api/proto
GEN_DIR := api/gen
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

generate:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)


build:
	go build -o bin/server ./cmd/server


.PHONY: test

test:
	go test -v ./...

.PHONY: all imports fmt lint

all: fmt

imports:
	@echo "Running goimports..."
	@goimports -w .
	@echo "Done."

fmt:
	@echo "Running gofmt..."
	@gofmt -w -s .
	@echo "Done."

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run
	@echo "Done."
