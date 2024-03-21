# Define the default tag for the Docker image
TAG ?= latest

# Define the version for the application
DEV_VERSION=4.4.0-dev
ENV=env GOOS=linux
TODAY:=$(shell date -u +%Y-%m-%dT%H:%M:%S)
TIMESTAMP:=$(shell date -u +%Y%m%d%H%M%S)
GITREV?=$(shell git rev-parse HEAD)
CELLS_VERSION?=${DEV_VERSION}.${TIMESTAMP}
GOBIN := /usr/local/go/bin/go


.PHONY: all clean build main dev docker docker-image build-docker-image dockercp start ds licenses

all: clean build

build: main

main:
	env CGO_ENABLED=0 ${GOBIN} build -a -trimpath \
		-ldflags "-X github.com/pydio/cells/v4/common.version=${CELLS_VERSION} \
		-X github.com/pydio/cells/v4/common.BuildStamp=${TODAY} \
		-X github.com/pydio/cells/v4/common.BuildRevision=${GITREV}" \
		-o cells .

dev:
	env CGO_ENABLED=0 ${GOBIN} build \
		-tags dev \
		-gcflags "all=-N -l" \
		-ldflags "-X github.com/pydio/cells/v4/common.version=${DEV_VERSION} \
		-X github.com/pydio/cells/v4/common.BuildStamp=2022-01-01T00:00:00 \
		-X github.com/pydio/cells/v4/common.BuildRevision=dev \
		-X github.com/pydio/cells/v4/common.LogFileDefaultValue=false \
		-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" \
		-o cells .

docker:
	env GOARCH=amd64 GOOS=linux ${GOBIN} build -trimpath \
		-ldflags "-X github.com/pydio/cells/v4/common.version=${CELLS_VERSION} \
		-X github.com/pydio/cells/v4/common.BuildStamp=${TODAY} \
		-X github.com/pydio/cells/v4/common.BuildRevision=${GITREV}" \
		-o cells-linux .

docker-image: docker
	docker build -t pydio .

build-docker-image: docker-image

dockercp:
	docker stop ${CONTAINER}; docker cp ./cells-linux ${CONTAINER}:/bin/cells; docker start ${CONTAINER}

start:
	./cells start

ds: dev start

licenses:
	go-licenses report . --template ${GOPATH}/src/github.com/google/go-licenses/testdata/modules/hello01/licenses.tpl > DEPENDENCIES

clean:
	rm -f cells cells-*

