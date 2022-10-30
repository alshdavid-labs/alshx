ifndef OUTDIR
	OUTDIR := ./dist
else
	OUTDIR := ${OUTDIR}
endif

ifndef RELEASE_DATE
	RELEASE_DATE := $(shell date -u --iso-8601=seconds)
else
	RELEASE_DATE := ${RELEASE_DATE}
endif

PROJECT_NAME := $(shell basename "$(CURDIR)")
VERSION := $(shell cat ./VERSION)
VERSION_FLAG := -X 'github.com/alshdavid/alshx/tools/$(PROJECT_NAME)/platform/meta.Version=$(VERSION)'
RELEASE_DATE_FLAG := -X 'github.com/alshdavid/alshx/tools/$(PROJECT_NAME)/platform/meta.ReleaseDate=$(RELEASE_DATE)'

ifeq (${PROD},true)
	LD_FLAGS := ${LD_FLAGS} -s -w $(VERSION_FLAG) $(RELEASE_DATE_FLAG)
else
	LD_FLAGS := ${LD_FLAGS} $(VERSION_FLAG) $(RELEASE_DATE_FLAG)
endif

ifeq (${GOOS},windows)
	BIN_NAME := $(PROJECT_NAME).exe
else ifeq ($(shell go env GOOS),windows)
	BIN_NAME := $(PROJECT_NAME).exe
else
	BIN_NAME := $(PROJECT_NAME)
endif

default: build

.PHONY: clean
clean:
	rm -r -f ./dist

.PHONY: build
build:
	go build -ldflags="$(LD_FLAGS)" -o "$(OUTDIR)/$(BIN_NAME)" ./cmd/main.go
