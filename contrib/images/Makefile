all: simd-env

simd-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag soominhyunwoosdk/simd-env simd-env

.PHONY: all simd-env
