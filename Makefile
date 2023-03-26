PROJECT := sb-ms-api
VERSION := 1.0.0-beta

HAS_BUILD_IMAGE := $(shell docker images | grep ${PROJECT} | grep -c build)

docker-build-if-nex:
ifeq ($(HAS_BUILD_IMAGE), 0)
	$(MAKE) docker-build-dev
else
	@echo "Build image exists"
endif


docker-build-dev:
	@echo "Building docker image..."
	@docker build --build-arg VERSION=${VERSION} \
		--build-arg APP_NAME=${PROJECT} \
		-t ${PROJECT}:build-prod \
		-f Dockerfile .

docker-run: docker-build-if-nex
	@echo "Starting docker containers..."
	@docker-compose -f docker-compose.yml up --build -d

docker-down:
	@docker-compose -f docker-compose.yml down

docker-logs: docker-build-if-nex
	@echo "Starting server in docker container..."
	@docker-compose -f docker-compose.yml logs

mock:
	@cd ./internal/mock

protogen:
	@cd ./pb; bash ./generate

cover:
	@go tool cover -html=coverage.out

test:
	@echo "Running tests..."
	@go test -timeout 30s ./... -short -count=1 -coverprofile=coverage.out

lint:
	@echo "Running golangci-lint..."
	@~/go/bin/golangci-lint run
