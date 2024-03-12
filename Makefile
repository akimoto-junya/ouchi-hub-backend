.PHONY: prepare
prepare:
	@docker build -t buf-connect ./docker/buf-connect

.PHONY: buf-generate
buf-generate:
	@docker run --rm --mount type=bind,src=./,dst=/work --workdir /work buf-connect generate

.PHONY: generate
generate: buf-generate

.PHONY: test
test:
	go test ./...
