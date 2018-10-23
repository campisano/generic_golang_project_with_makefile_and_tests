BINARY_NAME=calc
DEBUG_NAME=calc_debug
MAIN_PACKAGE=./cmd
SRC_DIRS?=$$(go list ./...)

.PHONY: all
all: check test

.PHONY: check
check: gometalinter testify
	@echo " -> Checking code..."
	@gometalinter.v2 --vendor --fast --sort=linter $(SRC_DIRS)

PKG_GOMETALINTER=gopkg.in/alecthomas/gometalinter.v2
.PHONY: gometalinter
gometalinter: $(GOPATH)/src/$(PKG_GOMETALINTER)
$(GOPATH)/src/$(PKG_GOMETALINTER):
	@echo " -> Installing Gometalinter..."
	@go get -u $(PKG_GOMETALINTER)
	@gometalinter.v2 --install

.PHONY: format
format:
	@echo " -> Formatting code..."
	@go fmt $(SRC_DIRS)

.PHONY: test
test:	export GOCACHE=off
test: testify build
	@echo " -> Testing code..."
	@go test -v $(SRC_DIRS)

PKG_TESTIFY=github.com/stretchr/testify
.PHONY: testify
testify: $(GOPATH)/src/$(PKG_TESTIFY)
$(GOPATH)/src/$(PKG_TESTIFY):
	@echo " -> Installing Testify..."
	@go get -u $(PKG_TESTIFY)

.PHONY: build
build:
	@echo " -> Building code..."
	@go build -v -o $(BINARY_NAME) $(MAIN_PACKAGE)

.PHONY: debug
debug:
	@echo " -> Building code for debug..."
	@go build -v -gcflags=all="-N -l" -o $(DEBUG_NAME) $(MAIN_PACKAGE)

.PHONY: clean
clean:
	@echo " -> Cleaning code..."
	@go clean -v $(SRC_DIRS)
	@rm -f $(BINARY_NAME) $(DEBUG_NAME)

.DEFAULT_GOAL := test
