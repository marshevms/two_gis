include config.mk

.PHONY: build
build:
	$(info building ...)
	go build -o $(BIN)/two_gis ./cmd/two_gis

.PHONY: run
run:
	$(info running ...)
	go run ./cmd/two_gis