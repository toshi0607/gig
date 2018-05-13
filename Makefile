PACKAGES = $(shell ./scripts/packages.sh)

EXTERNAL_TOOLS = \
    github.com/golang/dep/cmd/dep \
    github.com/laher/goxc \
    github.com/motemen/gobump \
    github.com/tcnksm/ghr \
    github.com/Songmu/ghch/cmd/ghch

setup:
	@for tool in $(EXTERNAL_TOOLS) ; do \
      echo "Installing $$tool" ; \
      go get $$tool; \
    done

test-all: vet lint test

test:
	./scripts/test.sh

vet:
	go vet ${PACKAGES}

lint:
	@go get github.com/golang/lint/golint
	golint ${PACKAGES} -set_exit_status

release: bump upload formula

bump: setup
	./scripts/bumpup.sh

upload: bump
	./scripts/upload.sh

formula: upload
	./scripts/formula.sh

.PHONY: test-all test vet lint setup release bump upload formula
