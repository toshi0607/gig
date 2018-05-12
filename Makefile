PACKAGES = $(shell ./scripts/packages.sh)

test-all: vet lint test

test:
	./scripts/test.sh

vet:
	go vet ${PACKAGES}

lint:
	@go get github.com/golang/lint/golint
	golint ${PACKAGES} -set_exit_status

.PHONY: test-all test vet lint
