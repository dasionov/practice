# Makefile for Go project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVENDOR=$(GOCMD) mod vendor

# Directories
BINDIR=bin
SRCDIR=cmd
PROTO_DIR=pkg/proto
PKGDIR=pkg
PACKAGEDIR=package
PACKAGEFILE=$(PACKAGEDIR)/my_project.tar.gz

# Linter
LINTER=golangci-lint

# Executable names
CLIENT_EXEC=$(BINDIR)/client
SERVER_EXEC=$(BINDIR)/server

# Protobuf
PROTOC=protoc
PROTO_SRC=$(PROTO_DIR)/*.proto
PROTO_GEN=$(PROTO_DIR)/*.pb.go

# Targets
all: build

build: proto $(CLIENT_EXEC) $(SERVER_EXEC)

$(CLIENT_EXEC): $(SRCDIR)/client/main.go
	$(GOBUILD) -o $(CLIENT_EXEC) ./$(SRCDIR)/client

$(SERVER_EXEC): $(SRCDIR)/server/main.go
	$(GOBUILD) -o $(SERVER_EXEC) ./$(SRCDIR)/server

lint:
	$(LINTER) run

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -rf $(BINDIR) $(PROTO_GEN) $(PACKAGEDIR)

vendor:
	$(GOVENDOR)

proto: $(PROTO_SRC)
	$(PROTOC) --go_out=$(PROTO_DIR) --go-grpc_out=$(PROTO_DIR) $(PROTO_SRC)

package: build
	mkdir -p $(PACKAGEDIR)
	tar -czvf $(PACKAGEFILE) $(BINDIR)

.PHONY: all build clean test lint vendor proto package
