.ONESHELL:
.PHONY: builder-shell build build-docker clean download fmt-check list-images run run-docker

PROJECT_NAME=calc-server-rust
BRANCH := $(shell git rev-parse --abbrev-ref HEAD | sed -e 's/-/./g;s=/=.=g')
COMMIT := $(shell git rev-parse --short HEAD)
APP_BIN=./app.bin.local
BUILD_BIN=target/release/rust_server
BUILDER_IMAGE=rust:1.85-slim-bookworm
BUILDER_CMD=cargo build --release  && ls -al target/*/rust_server
DOCKER_TAG=latest
RUNNER_IMAGE=debian:12

builder-shell:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		$(BUILDER_IMAGE) bash

fmt-check:
	rustfmt --check --edition 2021 src/main.rs

fmt:
	rustfmt --edition 2021 src/main.rs

download:
	cargo check

build:
	time $(BUILDER_CMD) \
	&& mv $(BUILD_BIN) $(APP_BIN) && ls -al $(APP_BIN)

build-by-docker:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		$(BUILDER_IMAGE) $(BUILDER_CMD) \
	&& mv $(BUILD_BIN) $(APP_BIN) && ls -al $(APP_BIN)

build-docker:
	docker build -t $(PROJECT_NAME):$(BRANCH) --progress plain \
		--build-arg CI_JOB_TOKEN=$(CI_JOB_TOKEN) \
		-f Dockerfile .
	docker images | grep $(PROJECT_NAME)

clean:
	docker container prune -f ; docker image prune -f ; docker volume prune -f
	rm -rf target $(APP_BIN)

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
