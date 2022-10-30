TOOLS_DIR := tools

build:
	$(foreach file, $(wildcard $(TOOLS_DIR)/*), echo $(file);)
