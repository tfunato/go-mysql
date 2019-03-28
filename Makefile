NAME     := go-mysql
VERSION  := v0.1
SRCS     := $(shell find . -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X main.version=$(VERSION)"

bin/$(NAME): $(SRCS) FORCE
	go build $(LDFLAGS) -o $@

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: FORCE
FORCE:
