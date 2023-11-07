BIN:=$(CURDIR)/bin


# linter
GOLANGCI:=$(BIN)/golangci-lint
GOLANGCI_VERSION:=1.55.1

.PHONY: golangci-install
golangci-install:
ifeq ($(wildcard $(GOLANGCI)),)
	$(info downloading golangci-lint)
	GOBIN=$(BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_VERSION)
endif

.PHONY: lint
lint: golangci-install
	$(GOLANGCI) run

.PHONY: test
test:
	go test ./...