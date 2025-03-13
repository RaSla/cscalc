.ONESHELL:
.PHONY: builder-shell build build-docker clean list-images run run-docker

PROJECT_NAME=calc-server-python
BRANCH := $(shell git rev-parse --abbrev-ref HEAD | sed -e 's/-/./g;s=/=.=g')
COMMIT := $(shell git rev-parse --short HEAD)
BUILDER_IMAGE=python:3.12-alpine
BUILDER_CMD=go build -v -o $(APP_BIN) ./src/main.go
RUNNER_IMAGE=python:3.12-alpine
RUNNER_CMD=python3 server.py

builder-shell:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		$(BUILDER_IMAGE) sh

download:
	pip3 install -r requirements.txt

build-docker:
	docker build -t $(PROJECT_NAME):$(BRANCH) --progress plain \
		-f Dockerfile .
	docker images | grep $(PROJECT_NAME)

clean:
	docker container prune -f ; docker image prune -f ; docker volume prune -f
	rm -rf .cache .config $(APP_BIN)

list-images:
	docker images | grep $(PROJECT_NAME)

run:
	$(RUNNER_CMD)

run-by-docker:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		-p 8080:8080 \
		$(RUNNER_IMAGE) $(RUNNER_CMD)

run-docker:
	docker run -it --rm -u `id -u`:`id -g` \
		-p 8080:8080 \
		$(PROJECT_NAME):$(BRANCH)
