VERSION := 3.0.0

DIST_DIR := bin
WEB_DIR := ${DIST_DIR}/web
WEB_WASM := $(WEB_DIR)/spaceinvaders.wasm
MODULE := github.com/programatta/spaceinvaders

.PHONY: build build-web run run-web clean

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o ${DIST_DIR}/spaceinvaders main.go

build-win:
	env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ${DIST_DIR}/spaceinvaders.exe main.go

# Requiere de un OSX para realizar compilación nativa con bindings de C
# build-mac:
# 	env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ${DIST_DIR}/spaceinvaders-mac main.go

# build-mac-arm:
# 	env GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION)" -o ${DIST_DIR}/spaceinvaders-macarm main.go

build-web:
	mkdir -p ${WEB_DIR}
	env GOOS=js GOARCH=wasm go build -ldflags "-X main.Version=$(VERSION)" -buildvcs=false -o ${WEB_WASM} ${MODULE}
	cp $$(go env GOROOT)/lib/wasm/wasm_exec.js ${WEB_DIR}
	printf '%s\n' \
	'<!DOCTYPE html>' \
	'<html>' \
	'  <head>' \
	'    <meta charset="UTF-8">' \
	'    <title>Space Invaders - Ebiten</title>' \
	'  </head>' \
	'  <body>' \
	'    <script src="wasm_exec.js"></script>' \
	'    <script>' \
	'      const go = new Go();' \
	'      WebAssembly.instantiateStreaming(fetch("spaceinvaders.wasm"), go.importObject).then(result => {' \
	'        go.run(result.instance);' \
	'      });' \
	'    </script>' \
	'  </body>' \
	'</html>' \
	> ${WEB_DIR}/index.html

build-all: build build-win build-web

run:
	go run main.go

run-web:
	go run github.com/hajimehoshi/wasmserve@latest .

clean:
	rm -rf ${DIST_DIR}
