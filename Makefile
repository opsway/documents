#!/usr/bin/env bash

# Where to push the docker image
REGISTRY?="quay.io/opsway"

# The binary to build
BIN := "$(shell basename "$(PWD)")"

IMAGE := "$(REGISTRY)/$(BIN)"

DATE := "$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")"
GIT_SHA1 := "$(shell git log -n 1 --pretty=format:%H)"

.DEFAULT_GOAL := build

t:
	@echo $(DATE)
	@echo $(GIT_SHA1)
fmt:
	 go fmt ./...

test: fmt
	 go test ./...

clean:
	rm -fr public/index.json
	rm -fr server

go-build: test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o server .

public/assets: # assets build
	mkdir -p public/assets
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/favicon-16x16.png --output public/assets/favicon-16x16.png
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/favicon-32x32.png --output public/assets/favicon-32x32.png
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/swagger-ui.css --output public/assets/swagger-ui.css
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/swagger-ui-bundle.js --output public/assets/swagger-ui-bundle.js
	curl https://raw.githubusercontent.com/swagger-api/swagger-ui/v3.20.6/dist/swagger-ui-standalone-preset.js --output public/assets/swagger-ui-standalone-preset.js

swagger-generate:
	swagger generate spec --output=public/index.json

docs: public/assets swagger-generate

start: go-build
	documents -addr 8080

build: clean go-build docs # image build
	docker build --build-arg BUILD_DATE="$(DATE)" --build-arg VCS_REF="$(GIT_SHA1)" --tag "$(IMAGE)" .

publish: build # image publish
	docker push "$(IMAGE)"
