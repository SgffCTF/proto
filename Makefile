PROTO_DIR := proto
PROTO_FILE := proto/*/*.proto

GEN_GO_DIR := gen/go
GEN_PYTHON_DIR := gen/python

install:
	@echo "Installing protoc-gen-go and protoc-gen-go-grpc..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go-generate:
	@echo "Generating Go code from $(PROTO_FILE)..."
	@protoc -I $(PROTO_DIR) $(PROTO_FILE) --go_out=$(GEN_GO_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GEN_GO_DIR) --go-grpc_opt=paths=source_relative

python-generate:
	@echo "Generating Python code from $(PROTO_FILE)..."
	@python -m grpc_tools.protoc -I=$(PROTO_DIR) --python_out=$(GEN_PYTHON_DIR) --grpc_python_out=$(GEN_PYTHON_DIR) $(PROTO_FILE)

all: install generate
