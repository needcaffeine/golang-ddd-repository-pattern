# Set a default target.
.DEFAULT_GOAL := help

PKG := ./cmd/golang-ddd-repository-pattern

# Output path for our binary.
BIN ?= .build/golang-ddd-repository-pattern

.PHONY: help # Show this help message.
help:
	@echo 'Usage: make [target] ...'
	@echo ''
	@echo 'Targets:'
	@fgrep -h "#!" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e "s/:.*#!/:/" | column -t -s":"
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1:\t\2/' | column -t -s":"

.PHONY: ---- # -----------------------
----:
	@echo "This is a dummy placeholder target. You can delete it."

.PHONY: all # clean, build, test.
all: clean build test

.PHONY: build # Build our binary (go build).
build:
	@echo "+ $@"
	$Qgo build -o $(BIN) $(PKG)

.PHONY: clean # Clean up binaries.
clean:
	@echo "+ $@"
	$Q$(RM) -fr .build

.PHONY: fmtchk # Run the formatter on every file.
fmtchk:
	@echo "+ $@"
	$Qexit $(shell gofmt -l . | grep -v '^vendor' | wc -l)

.PHONY: fmtfix # Run the formatter on every file, and fix all issues.
fmtfix:
	@echo "+ $@"
	@gofmt -w $(shell find . -iname '*.go' | grep -v vendor)

.PHONY: generate # Generate Go files by processing the source (go generate).
generate:
	@echo "+ $@"
	$Qgo generate ./...

.PHONY: run # Run the package (go run).
run:
	@echo "+ $@"
	$Qgo run $(PKG)

.PHONY: test # Run the `vet` and `fmtchk` targets, and generate code coverage reports.
test: vet fmtchk
	@CGO_ENABLED=1 go test -race -v ./... -coverpkg="./..." -coverprofile=.coverprofile
	@grep -v 'cmd' < .coverprofile > .covprof && mv .covprof .coverprofile
	@go tool cover -func=.coverprofile
	@go tool cover -html .coverprofile -o .coverprofile.html
	@open .coverprofile.html

.PHONY: vet # Vet is a tool that checks correctness of Go programs (go vet).
vet:
	@echo "+ $@"
	$Qgo vet ./...
