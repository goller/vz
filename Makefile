.PHONY: fmt test test-asan test-macos11 test-macos10

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

test-asan:
	@CGO_CFLAGS=-fsanitize=address CGO_LDFLAGS=-fsanitize=address go test -c . -o vz.test
	@codesign --entitlements ./example/linux/vz.entitlements -s - ./vz.test || true
	@./vz.test

test-macos10:
	@go test -tags=macos10 -c . -o vz.test
	@codesign --entitlements ./example/linux/vz.entitlements -s - ./vz.test || true
	@./vz.test
