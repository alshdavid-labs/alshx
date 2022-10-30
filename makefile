TOOLS_DIR := tools

ifndef GOOS
	GOOS := $(shell go env GOOS)
else
	GOOS := ${GOOS}
endif

ifndef GOARCH
	GOARCH := $(shell go env GOARCH)
else
	GOARCH := ${GOARCH}
endif

OUTDIR := $(CURDIR)/dist/$(GOOS)-$(GOARCH)

default: clean build

.PHONY: clean
clean:
	rm -r -f "$(OUTDIR)"

.PHONY: build
build:
	rm -r -f "$(OUTDIR)"
	$(foreach file, $(wildcard $(TOOLS_DIR)/*), \
		echo ===== ; \
		echo BUILDING $(file) @ $(GOOS)-$(GOARCH) ; \
		echo ----- ; \
		cd "$(CURDIR)/$(file)" && \
		env OUTDIR="$(OUTDIR)" make -s build; \
		echo ----- ; \
		echo COMPLETE ;\
		echo ===== ; \
		echo ;\
	)

.PHONY: upgrade
upgrade:
	$(foreach file, $(wildcard $(TOOLS_DIR)/*), \
		echo ===== ; \
		echo Upgrading $(file) ; \
		echo ----- ; \
		cd "$(CURDIR)/$(file)" && \
		go get -u ./... && \
		go mod tidy ; \
		echo ----- ; \
		echo COMPLETE ;\
		echo ===== ; \
		echo ;\
	)

