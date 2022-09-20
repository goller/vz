.PHONY: fmt test test-macos11
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
