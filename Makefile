.PHONY: test build amd64 arm64 arm i386

test:
	dep ensure
	go test ./...

build: .drone.sh
	./.drone.sh

amd64: Dockerfile
	docker build -t "plugins/hugo:amd64" --build-arg HUGO_VERSION="$(hugo)" --build-arg HUGO_ARCH=64bit .

arm64: Dockerfile.arm64
	docker build -t "plugins/hugo:arm64" --build-arg HUGO_VERSION="$(hugo)" --build-arg HUGO_ARCH=arm64 .

arm: Dockerfile.arm
	docker build -t "plugins/hugo:arm" --build-arg HUGO_VERSION="$(hugo)" --build-arg HUGO_ARCH=arm .