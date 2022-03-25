export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: build

ifdef VERSION
MAJOR := $(word 1, $(subst ., , $(VERSION)))
MINOR := $(word 2, $(subst ., , $(VERSION)))
PATCH := $(word 3, $(subst ., , $(VERSION)))
else
VERSION := $(shell git rev-list -1 HEAD)
endif
COMMIT := $(shell git rev-list -1 HEAD)

all: install build

install: # @HELP install the atomix command line tool
install:
	go install -ldflags "-X github.com/atomix/atomix-runtime/pkg/version.version=$(VERSION) -X github.com/atomix/atomix-runtime/pkg/version.commit=$(COMMIT)" ./cmd/atomix

build: # @HELP build the source code
build: deps
	GOOS=linux GOARCH=amd64 go build -gcflags=-trimpath=${GOPATH} -asmflags=-trimpath=${GOPATH} -o build/_output/atomix-runtime ./cmd/atomix-runtime

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...

test: # @HELP run the unit tests and source code validation
test: build license_check linters
	go test github.com/atomix/atomix-raft-storage/...

coverage: # @HELP generate unit test coverage data
coverage: build linters license_check

linters: # @HELP examines Go source code and reports coding problems
	GOGC=50 golangci-lint run

license_check: # @HELP examine and ensure license headers exist
	./build/bin/license-check

protos: # @HELP build Protobuf/gRPC generated types
protos:
	docker run -it -v `pwd`:/build \
		--entrypoint build/bin/compile_protos.sh \
		`docker build -q build/docker/protoc`

images: # @HELP build atomix storage controller Docker images
images: build
	docker build . -f build/atomix-runtime/Dockerfile -t atomix/atomix-runtime:latest
ifdef VERSION
	docker tag atomix/atomix-runtime:latest atomix/atomix-runtime:${MAJOR}.${MINOR}.${PATCH}
	docker tag atomix/atomix-runtime:latest atomix/atomix-runtime:${MAJOR}.${MINOR}
endif

kind: images
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image atomix/atomix-runtime:latest
ifdef VERSION
	kind load docker-image atomix/atomix-runtime:${MAJOR}.${MINOR}.${PATCH}
	kind load docker-image atomix/atomix-runtime:${MAJOR}.${MINOR}
endif

push: # @HELP push atomix-raft-node Docker image
push:
	docker push atomix/atomix-runtime:latest
ifdef VERSION
	docker push atomix/atomix-runtime:${MAJOR}.${MINOR}.${PATCH}
	docker push atomix/atomix-runtime:${MAJOR}.${MINOR}
endif

release: # @HELP release atomix-runtime Docker image
release: images push
