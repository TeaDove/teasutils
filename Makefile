GO ?= GO111MODULE=on CGO_ENABLED=0 go
VERSION ?= $(shell cat VERSION)

test:
	$(GO) test ./... -count=1 -p=100

update-all:
	$(GO) get -u ./...

tag:
	git tag $(VERSION)
	git push origin --tags

lint:
	gofumpt -w .
	golines --base-formatter=gofumpt --max-len=120 --no-reformat-tags -w .
	wsl --fix ./...
	golangci-lint run --fix