BINARY=mysql-toolkit

VERSION?=1.0.0
BUILD=`git rev-parse HEAD`
BUILD_TIME=`date +%FT%T%z`

PLATFORMS=darwin linux windows
ARCHITECTURES=amd64

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}  -X main.BuildTime=${BUILD_TIME}"

default: build

build:
	go build ${LDFLAGS} -o bin/${BINARY} ./main.go 

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES),\
	$(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o bin/$(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH) ./main.go )))
