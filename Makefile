CURDIR = $(shell pwd)
#GOPATH= $(dir $(abspath $(dir $(abspath $(dir ${CURDIR})))))
GOBIN = $(CURDIR)/build/bin
GO ?= latest
VERSION ?= undefined
OS ?= $(shell go env GOOS)
ARCH ?= $(shell go env GOARCH)
LDFLAGS = -s -w -X main.Version=$(VERSION)
ifeq (linux,$(OS))
	LDFLAGS+= -linkmode external -extldflags "-static"
endif

tool:
	@GOPATH=$(GOPATH) go build -v -o ./build/bin/palette-tool ./cmd/tools
	@echo "Done building."
	@echo "Run \"$(GOBIN)/palette-tool\" to launch palette-tool."

dist: clean
	@GOPATH=$(GOPATH) go build -ldflags='$(LDFLAGS)' -o ./build/bin/palette-tool ./cmd/tools
	@tar cfvz ./build/palette-tools_$(VERSION)_$(OS)_$(ARCH).tar.gz -C ./build/bin palette-tool
	@echo "Distribution file created."
	@ls -lh ./build

clean:
	rm -rf build
