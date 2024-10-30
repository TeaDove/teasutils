GO ?= GO111MODULE=on CGO_ENABLED=0 go

test:
	$(GO) test ./... -count=1 -p=100

update-all:
	$(GO) get -u ./...
