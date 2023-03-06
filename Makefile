.PHONY: all clean

BINARY_NAME=go-diydns
DIST_DIR=bin

all: clean build

build:
	go build -o $(DIST_DIR)/$(BINARY_NAME)

build-all:
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64
	GOOS=linux GOARCH=arm go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm
	GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm64
	GOOS=linux GOARCH=386 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-386
	GOOS=freebsd GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-freebsd-amd64
	GOOS=freebsd GOARCH=386 go build -o $(DIST_DIR)/$(BINARY_NAME)-freebsd-386
	GOOS=freebsd GOARCH=arm go build -o $(DIST_DIR)/$(BINARY_NAME)-freebsd-arm
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe
	GOOS=windows GOARCH=386 go build -o $(DIST_DIR)/$(BINARY_NAME)-windows-386.exe
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64

run:
	go build -o $(DIST_DIR)/$(BINARY_NAME)
	$(DIST_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(DIST_DIR)
	mkdir $(DIST_DIR)