BINARY=mysql-toolkit

VERSION?=0.1.0
BUILD?=`git rev-parse HEAD`
BUILD_TIME=`date -u +%FT%T`

PLATFORMS := linux/amd64 darwin/amd64
GOPLATFORM = $(subst /, ,$@)
GOOS = $(word 1, $(GOPLATFORM))
GOARCH = $(word 2, $(GOPLATFORM))

VERSION_RELEASE=v${VERSION} (${BUILD}) BuildTime: ${BUILD_TIME}
BINARY_RELEASE=$(BINARY)-$(VERSION)-$(GOOS)-$(GOARCH)

LDFLAGS=-ldflags "-X github.com/qorbani/mysql-toolkit/cmd.Version=${VERSION} -X github.com/qorbani/mysql-toolkit/cmd.Build=${BUILD}  -X github.com/qorbani/mysql-toolkit/cmd.BuildTime=${BUILD_TIME}"
SOURCES := $(shell find . -name '*.go' | grep -v vendor)

default: build

.PHONY: build
build: check
	@echo "[-] Building..."
	@echo "    - Version: ${VERSION_RELEASE}"
	@echo "    - Generate: ${BINARY}"
	@go build ${LDFLAGS} -o bin/${BINARY} ./main.go
	@echo "[+] Done" 

.PHONY: release
release: check $(PLATFORMS)
	@echo "[+] Done" 

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	@echo "[-] Building $@..."
	@echo "    - Version: v${VERSION} (${BUILD}) BuildTime: ${BUILD_TIME}"
	@echo "    - Binary: ${BINARY_RELEASE}"
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build ${LDFLAGS} -v -o bin/$(BINARY_RELEASE) ./main.go

.PHONY: check
check: check-vet check-fmt
	@echo "[+] check Done"

.PHONY: check-vet
check-vet:
	@echo "[-] Checking go vet..."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -ne 0 ]; then \
		echo "[x] Vet found suspicious constructs!"; \
		exit 1; \
	fi

.PHONY: check-fmt 
check-fmt: 
	@echo "[-] Checking gofmt..."
	$(eval FMT_FILES := $(shell gofmt -l $(SOURCES)))
	@if [[ -n "$(FMT_FILES)" ]]; then \
		echo "[x] gofmt needs running on the following files:"; \
		echo "    - $(FMT_FILES)"; \
		echo "    [?] Use \`make fmt\` to reformat code."; \
		exit 1; \
	fi

.PHONY: fmt
fmt:
	@echo "[-] Formating gofmt..."
	@gofmt -w $(SOURCES)
	@echo "[+] Done"

.PHONY: clean
clean:
	@echo "[-] Cleaning..."
	@rm -Rf ./bin/$(BINARY){,-*}
	@echo "[+] Done"