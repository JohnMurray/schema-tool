BUILD_OPTS=-p 4 -race
BIN_NAME=schema

default: test

test: compile
	@echo
	@echo "[running tests]"
	@go test .

compile: copyright
	go build $(BUILD_OPTS) -o $(BIN_NAME)
	go vet
	golint .
	@gotags -tag-relative=true -R=true -sort=true -f="tags" -fields=+l .

copyright:
	@echo "Applying copyright to all Go source files"
	@./script/copyright-header.sh

setup:
	go get -u github.com/tools/godep
	go get -u github.com/golang/lint/golint
	go get -u github.com/jstemmer/gotags
	godep restore
	npm install -g doctoc

doctoc:
	doctoc readme.md --github
