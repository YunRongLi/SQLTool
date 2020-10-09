GOPATH := $(CURDIR)/_build
export GOPATH
PATH := $(CURDIR)/_build/bin:$(PATH)
export PATH

REPO := $(shell git config remote.origin.url)

ifneq ($(REPO),)
GOPKG := $(shell python -c 'print("$(REPO)".replace("https://","").replace(".git",""))')
BIN := $(shell python -c 'print("$(GOPKG)".rsplit("/",1)[1])')
endif

SRC := *.go

LDFLAGS += -X "BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "BuildCommit=$(shell git rev-parse HEAD)"

BUILD_FLAGS = "-v"

.PHONY: all build prebuild clean

all: build

prebuild: $(SRC)
	go get -d -v ./...
	install -d $(GOPATH)/src/$(GOPKG)
	cp -a $(CURDIR)/*.go $(GOPATH)/src/$(GOPKG)

build: prebuild
	go build -o $(GOPATH)/bin/$(BIN)

build-only: $(SRC)
	go build -o $(GOPATH)/bin/$(BIN)

clean:
	@rm -rf $(GOPATH)/bin/$(BIN) $(GOPATH)/src/$(GOPKG)
