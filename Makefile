PROTO_DIR := proto
PROTO_FILE := proto/*/*.proto

GEN_DIR := gen/go

install:
	@echo "Installing protoc-gen-go and protoc-gen-go-grpc..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate: $(PROTO_FILE)
	@echo "Generating Go code from $(PROTO_FILE)..."
	@protoc -I $(PROTO_DIR) $(PROTO_FILE) --go_out=$(GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative

all: install generate