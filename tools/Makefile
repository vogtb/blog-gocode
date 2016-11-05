PWD := $(shell pwd)

OS := linux
ARCH := amd64
LDFLAGS := -X main.build=master

GOPATH := $(PWD)/.go:$(PWD)
GOENV := GOPATH=$(GOPATH)
GO := /usr/local/go/bin/go
GOBUILD := $(GOENV) GOOS=$(OS) GOARCH=$(ARCH) $(GO) build -i
GOIMPORTS := $(GOENV) PATH=$(PWD)/.go/bin goimports -w

DEPS := $(shell $(GOENV) $(GO) list -f '{{range .Imports}}{{.}} {{end}}' ./... | tr -s ' ' "\n" | grep -v "^libs" | sort | uniq)

SRC_DIR := ./src
TARGET_DIR := ./target


# common
.PHONY: .go deps fmt

.go:
	@echo "# Setting project GOPATH"
	mkdir -p $(PWD)/.go/bin $(PWD)/.go/pkg $(PWD)/.go/bin/src

deps: .go
	@echo "# Installing dependencies"
	@echo $(DEPS)
	@echo $(DEPS) | $(GOENV) xargs -n1 $(GO) get -d
	mkdir -p $(TARGET_DIR)


fmt: .go
	@echo "# Formatting the code"
	gofmt -w .

# commands
APPLICATIONS= myapp

.PHONY: $(APPLICATIONS)
$(APPLICATIONS): deps fmt
	@echo "# Building $@"
	$(GOBUILD) \
		-ldflags "$(LDFLAGS)" \
		-o $(TARGET_DIR)/$@.$(OS).$(ARCH) \
		apps/$@

# distribution

.PHONY: dist

dist:
	rm -fr $(TARGET_DIR)/*
	rm -fr $(PWD)/.go
	rm -fr $(PWD)/pkg/$(OS)_$(ARCH)*
