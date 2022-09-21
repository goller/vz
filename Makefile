.PHONY: fmt test test-macos11 test-macos10 all

all: all-arm64 all-amd64

fmt:
	@ls | grep -E '\.(h|m)$$' | xargs clang-format -i

test:
	@go test -c . -o vz.test
	@codesign --entitlements ./example/linux/vz.entitlements -s - ./vz.test || true
	@./vz.test

test-macos11:
	@go test -tags=macos11 -c . -o vz.test
	@codesign --entitlements ./example/linux/vz.entitlements -s - ./vz.test || true
	@./vz.test

test-macos10:
	@go test -tags=macos10 -c . -o vz.test
	@codesign --entitlements ./example/linux/vz.entitlements -s - ./vz.test || true
	@./vz.test

.PHONY: all-amd64 all-arm64 example-amd64 example-arm64
all-amd64 all-arm64: all-%: example-%
	CGO_ENABLED=1 GOARCH=$* go test -c . -o vz-macos12-$*.test
	CGO_ENABLED=1 GOARCH=$* go test -tags=macos11 -c . -o vz-macos11-$*.test
	CGO_ENABLED=1 GOARCH=$* go test -tags=macos10 -c . -o vz-macos10-$*.test

example-amd64:
	CGO_ENABLED=1 GOARCH=amd64 make -C example/linux

example-arm64:
	CGO_ENABLED=1 GOARCH=arm64 make -C example/linux
	CGO_ENABLED=1 GOARCH=arm64 make -C example/macOS
