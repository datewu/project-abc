#GO11MODULES=on
APP?=project-abc
REGISTRY?=gcr.io/images
VER?=0.0.1
COMMIT_SHA=$(shell git rev-parse --short HEAD)

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build  \
	-ldflags \
	"-s -w -X main.GitCommit=${COMMIT_SHA} -X main.SemVer=${VER}" \
	-o ${APP} 
	
dlv-debug: clean
	@echo "Building for delve debug..."
	@go build \
	-ldflags \
	"-s -w -X main.GitCommit=${COMMIT_SHA} -X main.SemVer=${VER}" \
	-gcflags="all=-N -l" \
	-o ${APP} 


.PHONY: run
## run: runs go run main.go
run:
	go run -race main.go

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@rm -rf ${APP}

.PHONY: test
## test: runs go test with default values
test:
	go test -timeout 300s -v -count=1 -race ./...


.PHONY: build-tokenizer
## build-tokenizer: build the tokenizer application
build-tokenizer:
	${MAKE} -c tokenizer build

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor
	
# helper rule for deployment
check-environment:
ifndef ENV
	$(error ENV not set, allowed values - `staging` or `production`)
endif

.PHONY: docker-build
## docker-build: builds the stringifier docker image to registry
docker-build: build
	docker build -t ${APP}:${COMMIT_SHA} .

.PHONY: docker-push
## docker-push: pushes the stringifier docker image to registry
docker-push: check-environment docker-build
	docker push ${REGISTRY}/${ENV}/${APP}:${COMMIT_SHA}

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'