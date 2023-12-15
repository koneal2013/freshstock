## Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
APP_NAME=freshstock
DOCKER_IMAGE_NAME=freshstock-api
DOCKER_CONTAINER_NAME=feshstock-api-container

# Output directory
OUT_DIR=bin

# Build directories
BUILD_DIRS=$(OUT_DIR)

.PHONY: all test clean

all: test compile

compile:
	mkdir -p $(BUILD_DIRS)
	$(GOBUILD) -o $(OUT_DIR)/$(APP_NAME) ./cmd

docker-build: test
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run: docker-build
	docker run --name $(DOCKER_CONTAINER_NAME) -d -it -p 8080:8080 $(DOCKER_IMAGE_NAME)

docker-down:
	docker stop $(DOCKER_CONTAINER_NAME) && docker rm $(DOCKER_CONTAINER_NAME)

clean:
	$(GOCLEAN)
	rm -rf $(OUT_DIR)

test:
	$(GOTEST) -v ./...

run: compile
	./bin/$(APP_NAME)
