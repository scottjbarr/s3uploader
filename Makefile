# See http://peter.bourgon.org/go-in-production/
GO ?= go

BUILD_DIR = build
DIST_DIR = dist

APP = s3worker
APP_BUILD = $(BUILD_DIR)/$(APP)

BUILD = $(GO) build

GO_FILES = `ls *.go | grep -v test | xargs echo`

all: clean build

build:
	mkdir -p $(BUILD_DIR)
	$(BUILD) -o $(APP_BUILD) cmd/s3worker/main.go

run:
	$(GO) run cmd/s3worker/main.go

clean:
	rm -rf $(BUILD_DIR)
