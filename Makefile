GOLANG_VERSION?=1.11.5
REGISTRY?=quay.io/opsway

BIN:=$(shell basename "$(PWD)")
REPO:=$(REGISTRY)/$(BIN)

DATE:=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_SHA1:=$(shell git log -n 1 --pretty=format:%H)
VERSION:=$(shell git describe --tags --always --dirty)

### These variables should not need tweaking.

PKGS:= $(shell go list ./...)

TAG:=$(VERSION)
IMAGE_BASE:=$(REPO):base
IMAGE_TEST:=$(REPO):test
IMAGE_PACKAGE:=$(REPO):$(TAG)

.DEFAULT_GOAL:=help

help: # Output usage documentation
	@echo "Usage: make <target>"
	@echo " "
	@echo "Commands:"
	@echo " "
	@@grep -E '^[a-z\-]+' $(MAKEFILE_LIST)
	@echo " "

fmt: # gofmt and goimports all go files
	 go fmt ./...

image-base: # build base image
	docker build \
		--tag $(IMAGE_BASE) \
		--file build/base/Dockerfile .

image-test: image-base # build test image
	docker build \
		--build-arg GOLANG_VERSION=$(GOLANG_VERSION) \
		--build-arg IMAGE_BASE=$(IMAGE_BASE) \
		--tag $(IMAGE_TEST) \
		--file build/test/Dockerfile .

image-package: docs # build package image
	docker build \
		--no-cache \
		--build-arg GOLANG_VERSION=$(GOLANG_VERSION) \
		--build-arg IMAGE_BASE=$(IMAGE_BASE) \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD_DATE=$(DATE) \
		--build-arg VCS_REF=$(GIT_SHA1) \
		--tag $(IMAGE_PACKAGE) \
		--file build/package/Dockerfile .

test-in-docker: image-test # run all tests
	docker run \
		--rm \
		--interactive \
		--tty \
		--volume "$(PWD):/src" \
		$(IMAGE_TEST) \
		make test

test: # run all tests
	@echo ">> TEST: coverage on"
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg, $(PKGS),\
	    echo -n "     ";\
		go test -coverprofile=coverage.out -covermode=atomic $(pkg) || exit 1;\
		tail -n +2 coverage.out >> coverage-all.out;)
	@go tool cover -html=coverage-all.out -o coverage-all.html

clean:
	rm -fr public/index.json
	rm -fr server

public/assets: # assets build
	mkdir -p public/assets
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/favicon-16x16.png --output public/assets/favicon-16x16.png
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/favicon-32x32.png --output public/assets/favicon-32x32.png
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/swagger-ui.css --output public/assets/swagger-ui.css
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/swagger-ui-bundle.js --output public/assets/swagger-ui-bundle.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/swagger-ui-standalone-preset.js --output public/assets/swagger-ui-standalone-preset.js

swagger-generate: public/assets
#	swagger generate spec --output=public/index.json
# FIXME #6

docs: public/assets swagger-generate

start:
	docker run \
		--rm \
		--publish 8515:8515 \
		$(IMAGE_PACKAGE)

build: image-package # image build

publish: build # image publish
	docker push $(IMAGE_PACKAGE)

say-image-name:
	@echo "image: $(IMAGE_PACKAGE)"

say-image-labels:
	@docker inspect \
		$(IMAGE_PACKAGE) \
		| jq .[0].Config.Labels

say-image-history:
	@docker history \
		$(IMAGE_PACKAGE)
