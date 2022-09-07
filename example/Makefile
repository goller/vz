.PHONY: all
all: build codesign

.PHONY: codesign
codesign:
	codesign --entitlements vz.entitlements -s - ./virtualization

.PHONY: build cross-build
build:
	CGO_CFLAGS=-mmacosx-version-min=11.0 go build -o virtualization .

cross-build: virtualization-amd64 virtualization-arm64

virtualization-amd64 virtualization-arm64: virtualization-%: force-build
	CGO_CFLAGS=-mmacosx-version-min=11.0 CGO_ENABLED=1 GOOS=darwin GOARCH=$* go build -o $@ .

# the go compiler is doing a good job at not rebuilding unchanged files
# this phony target ensures out/vfkit-* are always considered out of date
# and rebuilt. If the code was unchanged, go won't rebuild anything so that's
# fast. Forcing the rebuild ensure we rebuild when needed, ie when the source code
# changed, without adding explicit dependencies to the go files/go.mod
.PHONY: force-build
force-build:
