ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	cd lib/zinctx && cargo build --release
	cp lib/zinctx/target/release/libzinctx.dylib lib/
	echo 'ROOT_DIR is $(ROOT_DIR)'
	go build -ldflags="-r $(ROOT_DIR)lib" main.go

run: build
	./main