PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test: lint
	go test $(PKGS)

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

$(GOMETALINTER):
	#go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

.PHONY: lint
lint: $(GOMETALINTER)
	gometalinter ./... --vendor --errors

BINARY := BuyingcreditBackend
VERSION ?= vlatest
PLATFORMS := windows linux
ARCHS := amd64 arm
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=$(ARCHS) go build -o release/$(BINARY)-$(VERSION)-$(os)

.PHONY: release
release: windows linux
