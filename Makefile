PACKAGES = $(shell ./scripts/packages.sh)

test-all: vet lint test

test:
	go test -v -parallel=4 ${PACKAGES} -race -coverprofile=profile.out -covermode=atomic

vet:
	go vet ${PACKAGES}

lint:
	@go get github.com/golang/lint/golint
	go list ./... | grep -v vendor | xargs -n1 golint -set_exit_status

.PHONY: test-all test vet lint
