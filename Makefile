PROJECT = golang-project-template
PROJECT_DIR = $(GOPATH)src

all: fmt prepare build

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

prepare:
	mkdir -p bin

# Build binarys
build:
	go build -o bin/golang-project-template cmd/main.go

# Build binarys
clean:
	rm -rf bin

.PHONY: copy-golang-project-template
golang-project-template:
	@mkdir -p $(PROJECT_DIR)
	@cp -rf $(shell command pwd;) $(PROJECT_DIR)
	@mkdir -p $(PROJECT_DIR)/$(PROJECT)/bin
	@cd $(PROJECT_DIR)/$(PROJECT)
	@echo $(PROJECT_DIR)/$(PROJECT)

.PHONY: test-style
test-style:
	@scripts/teststyle.sh

.PHONY: test-unit
test-unit:
	@scripts/testunit.sh

.PHONY: coverage
coverage:
	@scripts/coverage.sh
