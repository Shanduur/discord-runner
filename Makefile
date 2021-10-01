TARGET=discord-runner

.PHONY: build
build:
	go build -o ./build/$(TARGET) ./main.go

.PHONY: test
test:
	go test ./...

.PHONY: docker
docker:
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output=type=registry \
		--no-cache \
		--tag shanduur/$(TARGET):$(cat ./version.txt) \
		--tag shanduur/$(TARGET):dev \
		.

include .dev.env
export

.PHONY: run
run: build
	./build/$(TARGET)