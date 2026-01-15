VERSION ?= $(shell cat VERSION)

tag:
	git tag utils/$(VERSION)
	git push origin --tags

test:
	gotestsum --format-hide-empty-pkg -- ./... --race

lint:
	gofumpt -w .
	golines --base-formatter=gofumpt --max-len=120 --no-reformat-tags -w .
	wsl --fix ./...
	golangci-lint run --fix

test-and-lint: test lint