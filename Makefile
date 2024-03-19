.PHONY: prepare
prepare:
	@docker build -t buf-connect ./docker/buf-connect

.PHONY: buf-generate
buf-generate:
	@docker run --rm --mount type=bind,src=./,dst=/work --workdir /work buf-connect generate

.PHONY: generate
generate: buf-generate

.PHONY: buf-lint
buf-lint:
	@docker run --rm --mount type=bind,src=./,dst=/work --workdir /work buf-connect lint

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	@find ./ -type f \( -name "*.pb.go" -o -name "*.connect.go" \) -exec rm {} +
