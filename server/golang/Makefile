.ONESHELL:
.PHONY: builder-shell build build-docker clean download fmt-check list-images run run-docker

PROJECT_NAME=calc-server-golang
BRANCH := $(shell git rev-parse --abbrev-ref HEAD | sed -e 's/-/./g;s=/=.=g')
COMMIT := $(shell git rev-parse --short HEAD)
APP_BIN=./app.bin.local
BUILDER_IMAGE=golang:1.23.3-alpine
BUILDER_CMD=go build -v -o $(APP_BIN) ./src/main.go
#BUILDER_CMD=go build -v -mod readonly -o $(APP_BIN) ./cmd
DOCKER_TAG=latest
DP_GO_LINT_IMAGE=docker/golangci-lint:1.58.0-go1.22.3
RUNNER_IMAGE=debian:12

builder-shell:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		-e GOSUMDB="off" \
		$(BUILDER_IMAGE) sh

lint:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		-e GOSUMDB="off" \
		$(DP_GO_LINT_IMAGE) golangci-lint run -v

fmt-check:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		$(BUILDER_IMAGE) gofmt -d ./

fmt:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		$(BUILDER_IMAGE) gofmt -w ./

download:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		-e GOSUMDB="off" \
		$(BUILDER_IMAGE) go mod download -x

build:
	$(BUILDER_CMD) \
	&& chmod u+x ./app.* && ls -al app.*

build-by-docker:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		-e GOSUMDB="off" \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		$(BUILDER_IMAGE) $(BUILDER_CMD) \
	&& chmod u+x ./app.* && ls -al app.*

build-docker:
	docker build -t $(PROJECT_NAME):$(BRANCH) --progress plain \
		--build-arg CI_JOB_TOKEN=$(CI_JOB_TOKEN) \
		-f Dockerfile .
	docker images | grep $(PROJECT_NAME)

clean:
	docker container prune -f ; docker image prune -f ; docker volume prune -f
	rm -rf .cache .config $(APP_BIN)

list-images:
	docker images | grep $(PROJECT_NAME)

run:
	ls -al $(APP_BIN) \
	&& $(APP_BIN)

run-by-docker:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		-p 8080:8080 \
		$(RUNNER_IMAGE) $(APP_BIN)

run-docker:
	docker run -it --rm -u `id -u`:`id -g` \
		-p 8080:8080 \
		$(PROJECT_NAME):$(BRANCH)
