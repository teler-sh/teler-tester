.DEFAULT_GOAL := build

GO_MOD_VERSION := $(shell grep -Po '^go \K([0-9]+\.[0-9]+(\.[0-9]+)?)$$' go.mod)
GO := go${GO_MOD_VERSION}
ifeq ($(shell which ${GO}),)
	GO = go
endif

TELER_WAF_VERSION := $(shell ${GO} list -m -f '{{ .Version }}' github.com/teler-sh/teler-waf | cut -c 2-)

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy

verify:
	$(GO) mod verify

teler-waf-version:
	@echo "${TELER_WAF_VERSION}"

teler-waf:
	$(GO) get -v github.com/teler-sh/teler-waf@latest

build: tidy verify vet # teler-waf
	$(GO) build -ldflags "-s -w -X 'main.telerWAFVersion=${TELER_WAF_VERSION}'" -o demo .