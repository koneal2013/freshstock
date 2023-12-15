## Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
APP_NAME=freshstock
DOCKER_IMAGE_NAME=freshstock-api

# Output directory
OUT_DIR=bin

# Build directories
BUILD_DIRS=$(OUT_DIR)

.PHONY: all test clean

all: test compile

compile:
	mkdir -p $(BUILD_DIRS)
	$(GOBUILD) -o $(OUT_DIR)/$(APP_NAME) .

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run -it -p 8080:8080 $(DOCKER_IMAGE_NAME)

clean:
	$(GOCLEAN)
	rm -rf $(OUT_DIR)

test:
	$(GOTEST) -v ./...