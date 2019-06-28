go      ?= CGO_ENABLED=0 go
targets ?= srv

.PHONY: all
all: build

.PHONY: build
build: ${targets}

${targets}: ${targets}.go
	${go} build -o $@ ${@}.go

.PHONY: test
test:
	go vet

.PHONY: clean
clean:
	rm -f ${targets}
