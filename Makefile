.PHONY: \
	buf-breaking \
	buf-lint \
	build-descriptor \
	clean \
	gen

buf-breaking:
	@buf breaking --against .git#branch=origin/main

buf-lint:
	@buf lint

buf-format:
	@buf format -w

build-descriptor: gen
	@buf build -o gen/coffeeco.bin

gen: buf-lint
	@buf generate

clean:
	rm -rf gen/
