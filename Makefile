# Makefile for data-linter
# Targets:
#       test: Run the tests

all: build

test: test_unit

test_unit:
	go test -v -cover github.com/GabbyyLS/data-linter/linter/...
