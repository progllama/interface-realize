.PHONY: build

build:
	go build -o ./bin/interface-realize main.go

.PHONY: test

test:
	go test -cover -v -bench . ./...