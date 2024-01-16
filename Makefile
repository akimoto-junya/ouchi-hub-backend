GCP_LOCATION := asia-northeast1
GCP_PROJECT_ID := ouchihub
GCP_REPO_URL := $(GCP_LOCATION)-docker.pkg.dev/$(GCP_PROJECT_ID)/ouchi-hub-backend/api
VERSION := $(shell git rev-parse HEAD)

.PHONY: docker-build
docker-build:
	docker build -t $(GCP_REPO_URL):$(VERSION) .

.PHONY: docker-push
docker-push:
	docker push $(GCP_REPO_URL):$(VERSION)
