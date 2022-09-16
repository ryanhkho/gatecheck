BIN="./bin"
SRC=$(shell find . -name "*.go")

ifeq (, $(shell which richgo))
$(warning "could not find richgo in $(PATH), run: go install github.com/kyoh86/richgo@latest")
endif

.PHONY: fmt test install_deps clean

default: all

all: fmt test

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

test: install_deps
	$(info ******************** running tests ********************)
	richgo test -v ./...

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

clean:
	rm -rf $(BIN)
