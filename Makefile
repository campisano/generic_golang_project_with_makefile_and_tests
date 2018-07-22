GO=go
BINARY_NAME=calc
SRC_EXT=.go
SRC_DIRS=./
SRC_FILES := $(shell find $(SRC_DIRS) -name \*$(SRC_EXT))

.PHONY: all
all: check test

.PHONY: check
check: gometalinter testify
	@echo " -> Checking code..."
	@gometalinter.v2 --vendor --fast --sort=linter ./...

PKG_GOMETALINTER=gopkg.in/alecthomas/gometalinter.v2
.PHONY: gometalinter
gometalinter: $(GOPATH)/src/$(PKG_GOMETALINTER)
$(GOPATH)/src/$(PKG_GOMETALINTER):
	@echo " -> Installing Gometalinter..."
	@$(GO) get -u $(PKG_GOMETALINTER)
	@gometalinter.v2 --install

.PHONY: test
test: testify build
	@echo " -> Testing code..."
	@$(GO) test -v ./...

PKG_TESTIFY=github.com/stretchr/testify
.PHONY: testify
testify: $(GOPATH)/src/$(PKG_TESTIFY)
$(GOPATH)/src/$(PKG_TESTIFY):
	@echo " -> Installing Testify..."
	@$(GO) get -u $(PKG_TESTIFY)

.PHONY: run
run: build
	@echo " -> Running code..."
	@./$(BINARY_NAME)

.PHONY: build
build: $(BINARY_NAME)
$(BINARY_NAME): $(SRC_FILES)
	@echo " -> Building code..."
	@$(GO) build -o $(BINARY_NAME) -v

.PHONY: clean
clean:
	@echo " -> Cleaning code..."
	@$(GO) clean
	@rm -f $(BINARY_NAME)
