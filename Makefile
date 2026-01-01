FUZZ_TARGETS = Add Sub Mul Div ArithmeticUint64 ConvertSignedToInt8 ConvertSignedToUnsigned ConvertUnsignedToSigned ConvertUnsignedToUnsignedSmall

.PHONY: all tests test test-examples bench fuzz

all: tests bench fuzz
tests: test test-examples

test:
	go test -v -run ^Test -race ./...

test-examples:
	go test -v -run ^Example ./...

bench:
	go test -run - -benchmem -bench . ./...

fuzz:
	@for target in $(FUZZ_TARGETS); do \
		echo "Fuzzing $$target..."; \
		go test -run - -fuzz "Fuzz$$target" -fuzztime 1m ./...; \
	done
