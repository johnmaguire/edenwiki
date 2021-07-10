test:
	go test $(shell go list ./...)

bin:
	go build -trimpath -o ./edenwiki ./cmd/edenwiki

.PHONY: bin
.DEFAULT_GOAL := bin
