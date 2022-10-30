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

clean:
	rm -r -f ./dist

build:
	rm -r -f "$(OUTDIR)"

	$(foreach file, $(wildcard $(TOOLS_DIR)/*), \
		echo ===== ; \
		echo BUILDING $(file) ; \
		echo ----- ; \
		cd "$(CURDIR)/$(file)" && \
		env OUTDIR="$(OUTDIR)" make -s build; \
		echo ----- ; \
		echo COMPLETE ;\
		echo ===== ; \
		echo ;\
	)
