.PHONY: test build amd64 arm64 arm i386

test:
	dep ensure
	go test ./...

build: .drone.sh
	./.drone.sh

amd64: Dockerfile
	docker build -t "plugins/hugo:amd64" .

arm64: Dockerfile.arm64
	docker build -t "plugins/hugo:arm64" .

arm: Dockerfile.arm
	docker build -t "plugins/hugo:arm" .

arm: Dockerfile.i386
	docker build -t "plugins/hugo:i386" .