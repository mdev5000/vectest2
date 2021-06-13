
clean: FORCE
	rm -rf public/wasm
	rm -rf .tmpbins
	# @todo add removal of .vhtml.go files.

build: gen.frontend build.frontend FORCE

# Install dependencies
deps:
	go install github.com/mdev5000/runnr/cmd/runnr@latest
	go install github.com/mdev5000/tvecty/cmd/tvecty@latest

deps.proto:
	go get -u github.com/gogo/protobuf/protoc-gen-gofast

gen.frontend: FORCE
	go generate ./webf

build.backend: FORCE
	rm -rf build
	mkdir -p build
	go build -o build/server cmd/server/main.go

build.frontend: FORCE
	mkdir -p public/wasm
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" public/wasm/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o public/wasm/main.wasm ./cmd/frontend/main.go

proto: FORCE
	protoc proto/*.proto \
		--gofast_out=plugins=grpc:. \
		--proto_path=$(go list -f '{{ .Dir }}' -m github.com/gogo/protobuf) \
		--proto_path=.

update.tvecty: FORCE
	go install github.com/mdev5000/tvecty/cmd/tvecty@latest

html.convert: FORCE
	cat conv.html | vectyhtmlconv > conv.result.txt

FORCE: