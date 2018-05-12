setup:
	go get -v -u \
		github.com/golang/dep/cmd/dep \
		github.com/laher/goxc \
		github.com/tcnksm/ghr

deps: setup
	dep ensure

test: deps
	go test -race -cover -v ./...

.PHONY: setup deps test