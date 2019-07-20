GOLANG_VERSION?=1.11.5
SWAGGER_UI_VERSION?=3.20.6
SWAGGER_UI_DIST_URL?=https://raw.githubusercontent.com/swagger-api/swagger-ui/v$(SWAGGER_UI_VERSION)/dist
REGISTRY?=quay.io/opsway

BIN:=$(shell basename "$(PWD)")
REPO:=$(REGISTRY)/$(BIN)

SHELL:=/bin/bash
DATE:=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VCS_REF:=$(shell git log -n 1 --pretty=format:%H)
VERSION:=$(shell git describe --tags --always --dirty)

DOCKER_CMD:=docker run --rm --interactive --tty --volume $(PWD):/src

WKHTML_IMAGE:=$(REPO):wkhtml
WKHTML_IMAGE_FILE:=build/wkhtml.docker

IMAGE_BASE:=$(WKHTML_IMAGE)
BASE_IMAGE:=$(WKHTML_IMAGE)

DEVELOP_IMAGE:=$(REPO):develop
DEVELOP_IMAGE_FILE:=build/develop.docker

RELEASE_IMAGE:=$(REPO):$(VERSION)
RELEASE_IMAGE_ALIASE:=$(REPO):latest

.DEFAULT_GOAL:=help

help: # Output usage documentation
	@echo "Usage: make <target>"
	@echo " "
	@echo "Commands:"
	@echo " "
	@@grep -E '^[a-z\-]+' $(MAKEFILE_LIST)
	@echo " "

image-wkhtml: # build wkhtml image
	docker build \
		--build-arg BUILD_DATE=$(DATE) \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_REF=$(VCS_REF) \
		--tag $(WKHTML_IMAGE) \
		--file $(WKHTML_IMAGE_FILE) . |& tee image-wkhtml

image-develop: image-wkhtml # build develop image
	docker build \
		--build-arg GOLANG_VERSION=$(GOLANG_VERSION) \
		--build-arg IMAGE_BASE=$(IMAGE_BASE) \
		--build-arg BUILD_DATE=$(DATE) \
		--build-arg VERSION=$(VERSION) \
		--build-arg VCS_REF=$(VCS_REF) \
		--tag $(DEVELOP_IMAGE) \
		--file $(DEVELOP_IMAGE_FILE) . |& tee image-develop

image-release: image-develop # build release image
	docker build \
		--no-cache \
		--build-arg DEVELOP_IMAGE=$(DEVELOP_IMAGE) \
		--build-arg BASE_IMAGE=$(BASE_IMAGE) \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD_DATE=$(DATE) \
		--build-arg VCS_REF=$(VCS_REF) \
		--tag $(RELEASE_IMAGE) \
		--tag $(RELEASE_IMAGE_ALIASE) \
		--file build/release.docker .

run-in-docker: image-develop # run command in docker, use: cmd=<command>
	$(DOCKER_CMD) $(DEVELOP_IMAGE) $(cmd)

fmt: # gofmt and goimports all go files
	go fmt ./...

lint:
	goreportcard-cli -v -t 100

codequality: lint test

entrypoint: codequality docs
	CGO_ENABLED=0 go build  \
		-o entrypoint \
		main.go

test: # run all tests
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out

clean:
	rm -fr public/assets
	rm -fr public/index.json
	rm -fr entrypoint

public/assets: # assets build
	mkdir -p public/assets
	curl $(SWAGGER_UI_DIST_URL)/favicon-16x16.png --output public/assets/favicon-16x16.png
	curl $(SWAGGER_UI_DIST_URL)/favicon-32x32.png --output public/assets/favicon-32x32.png
	curl $(SWAGGER_UI_DIST_URL)/swagger-ui.css --output public/assets/swagger-ui.css
	curl $(SWAGGER_UI_DIST_URL)/swagger-ui-bundle.js --output public/assets/swagger-ui-bundle.js
	curl $(SWAGGER_UI_DIST_URL)/swagger-ui-standalone-preset.js --output public/assets/swagger-ui-standalone-preset.js

public/index.json: public/assets
	swagger generate spec --output=public/index.json

docs: clean public/assets public/index.json

start:
	docker run \
		--rm \
		--publish 8515:8515 \
		$(RELEASE_IMAGE)

build: entrypoint

publish-release: image-release # image publish release image
	docker push $(RELEASE_IMAGE)
	docker push $(RELEASE_IMAGE_ALIASE)

say-image-name:
	@echo "image: $(RELEASE_IMAGE)"

say-image-labels:
	@docker inspect \
		$(RELEASE_IMAGE) \
		| jq .[0].Config.Labels

say-image-history:
	@docker history \
		$(RELEASE_IMAGE)
