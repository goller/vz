.PHONY: all
all: build codesign

.PHONY: codesign
codesign:
	codesign --entitlements vz.entitlements -s - ./virtualization

.PHONY: build
build:
	CGO_CFLAGS=-mmacosx-version-min=11.0 go build -o virtualization .
