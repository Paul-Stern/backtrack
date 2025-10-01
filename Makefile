APP_NAME=backtrack
PLOG_PATH=github.com/rs/zerolog/cmd/prettylog
VERSION=$(shell git describe --abbrev=0 --tags)
BUILD_TIME=$(shell date +%Y-%m-%d_%H:%M:%S)
COMMIT_HASH=$(shell git rev-parse --short HEAD)
ARC_NAME=$(APP_NAME)-$(VERSION)-$(COMMIT_HASH)
GO_PACKAGE_PATH=$(shell go list -m)

.PHONY: build

build:
	@echo "Building $(APP_NAME) version $(VERSION)"
	go build  -ldflags="\
		-X '$(GO_PACKAGE_PATH)/version.Version=$(VERSION)' \
		-X '$(GO_PACKAGE_PATH)/version.Commit=$(COMMIT_HASH)' \
		-X '$(GO_PACKAGE_PATH)/version.BuildTime=$(BUILD_TIME)'" \
		-o $(APP_NAME) ./cmd/$(APP_NAME)
	go build -o prettylog $(PLOG_PATH)

bin-version: build
	@echo "$(APP_NAME) --version"
	./$(APP_NAME) --version

# For production builds
release: build
	@echo "Creating release $(VERSION)-$(COMMIT_HASH)"
	tar czf $(ARC_NAME).tar.gz $(APP_NAME) prettylog

# For development (without version info)
dev:
	go build ./cmd/$(APP_NAME) ./cmd/$(LOG_APP_NAME)

clean:
	rm -f $(APP_NAME) $(LOG_APP_NAME) *.tar.gz

.PHONY: version
version:
	@echo "Version: $(VERSION)"
	@echo "Commit: $(COMMIT_HASH)"
	@echo "BuildTime: $(BUILD_TIME)"
