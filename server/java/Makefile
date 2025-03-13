.ONESHELL:
.PHONY: builder-shell build build-docker clean list-images run run-docker

PROJECT_NAME=calc-server-java
BRANCH := $(shell git rev-parse --abbrev-ref HEAD | sed -e 's/-/./g;s=/=.=g')
COMMIT := $(shell git rev-parse --short HEAD)
BUILDER_IMAGE=amazoncorretto:17-alpine
#BUILDER_IMAGE=openjdk:11-jdk-slim
BUILDER_CMD=sh -c "javac -d ./bin src/com/company/Main.java src/com/company/Server.java"
#BUILDER_CMD=javac -d . src/test_java.java
RUNNER_IMAGE=amazoncorretto:17
RUNNER_CMD=java -classpath ./bin com.company.Main
#RUNNER_CMD=sh -c "time java -classpath . bin.test_java"

builder-shell:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		$(BUILDER_IMAGE) sh

build:
	$(BUILDER_CMD) \
	&& ls -al bin/com/company/

build-by-docker:
	docker run -it --rm -u `id -u`:`id -g` -v `pwd`:/mnt -w /mnt \
		-e HOME="/mnt" \
		$(BUILDER_IMAGE) $(BUILDER_CMD) \
	&& ls -al bin/com/company/

build-docker:
	docker build -t $(PROJECT_NAME):$(BRANCH) --progress plain \
		-f Dockerfile .
	docker images | grep $(PROJECT_NAME)

clean:
	docker container prune -f ; docker image prune -f ; docker volume prune -f
	rm -rf com

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
