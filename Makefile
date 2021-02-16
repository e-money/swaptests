build:
	go build -o swap

lint:
	golangci-lint run

# go get mvdan.cc/gofumpt
fmt:
	gofumpt -w **/*.go

# go get go get github.com/daixiang0/gci
imp:
	gci -w **/*.go

.PHONY: build fmt lint imp
