test:
	go test $(shell go list ./...)

bin:
	go build -trimpath -o ./gardenwiki ./cmd/gardenwiki

.PHONY: bin
.DEFAULT_GOAL := bin
