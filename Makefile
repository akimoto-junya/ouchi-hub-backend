.PHONY: prepare
prepare:
	@docker build -t buf-connect ./docker/buf-connect

.PHONY: buf-generate
buf-generate:
	@docker run --rm --mount type=bind,src=./,dst=/work --workdir /work buf-connect generate

.PHONY: sqlc-generate
sqlc-generate:
	@docker run --rm --mount type=bind,src=./,dst=/work --workdir /work sqlc/sqlc generate

.PHONY: generate
generate: buf-generate sqlc-generate
	go generate ./...

.PHONY: buf-lint
buf-lint:
	@docker run --rm --mount type=bind,src=./,dst=/work --workdir /work buf-connect lint

.PHONY: sqlc-vet
sqlc-vet:
	@docker run --rm --mount type=bind,src=./,dst=/work --workdir /work sqlc/sqlc vet

.PHONY: lint
lint: buf-lint sqlc-vet
	golangci-lint run ./...

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	@find ./ -type f \( -name "*.pb.go" -o -name "*.connect.go" \) -exec rm {} +
