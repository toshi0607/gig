PACKAGES = $(shell go list ./... | grep -v '/vendor/')

test-all: vet lint test

test:
	go test -v -parallel=4 ${PACKAGES}

vet:
	go vet ${PACKAGES}

lint:
	@go get github.com/golang/lint/golint
	go list ./... | grep -v vendor | xargs -n1 golint -set_exit_status

.PHONY: test-all test vet lint
