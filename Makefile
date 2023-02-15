.PHONY: build.filters
build.filters:
	@find ./filters -mindepth 1 -type f -name "main.go" \
	| xargs -I {} bash -c 'dirname {}' \
	| xargs -I {} bash -c 'cd {} && tinygo build -o main.wasm -scheduler=none -target=wasi ./main.go'

