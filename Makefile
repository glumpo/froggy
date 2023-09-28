# Set VERSION env variable in form x.y.z to provide version.

GIT_COMMIT=`git rev-parse HEAD`
BUILD_DATE=`date +%d-%m-%Y-%H:%M`
VERSION?=0.0.0-develop

GOOS?=linux
GOARCH?=amd64

all: froggy

.PHONY: froggy
froggy:
	@echo "Version: $(VERSION)"
	@echo "Commit: $(GIT_COMMIT)"
	@echo "Date: $(BUILD_DATE)"
	@echo "GOOS: $(GOOS)"
	@echo "GOARCH: $(GOARCH)"

	mkdir -p ./dist

	go get ./...
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(GIT_COMMIT) -X main.date=$(BUILD_DATE)" -o ./dist/froggy ./cmd/froggy/main.go

	@echo "Done"

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 make froggy

.PHONY: old_darwin
old_darwin:
	GOOS=darwin GOARCH=amd64 make froggy

.PHONY: darwin
darwin:
	GOOS=darwin GOARCH=arm64 make froggy
