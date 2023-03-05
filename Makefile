.PHONY: all clean

BINARY_NAME=go-diydns
DIST_DIR=bin

all: clean build

build:
	go build -o $(DIST_DIR)/$(BINARY_NAME)

run:
	go build -o $(DIST_DIR)/$(BINARY_NAME)
	$(DIST_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(DIST_DIR)
	mkdir $(DIST_DIR)