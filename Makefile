lint:
	golangci-lint run --exclude-use-default=false --enable=golint --enable=goimports --enable=unconvert --enable=unparam --enable=gosec

test:
	go test -v .

.PHONY: lint test