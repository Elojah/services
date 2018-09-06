PACKAGE  = launcher
DATE    ?= $(shell date +%FT%T%z)
VERSION  ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))

GO      = go1.11
GODOC   = godoc
GOFMT   = gofmt
GOLINT  = gometalinter

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[0;35m▶\033[0m")

.PHONY: all
all: check

# Vendor
.PHONY: vendor
vendor:
	$(info $(M) running go mod vendor…) @
	$Q $(GO) mod vendor
	# $Q modvendor -copy="**/*.c **/*.h" -v

# Check
.PHONY: check
check: lint test

# Lint
.PHONY: lint
lint: $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q go get github.com/golang/lint/golint
	$Q gometalinter "--vendor" \
					"--disable=gotype" \
					"--fast" \
					"--json" \
					"./..." \

# Test
.PHONY: test
test:
	go test -cover -race -v ./...

.PHONY: fmt
fmt:
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc:
	$(info $(M) running $(GODOC)…) @
	$Q $(GODOC) ./...

.PHONY: version
version:
	@echo $(VERSION)
