.PHONY: build
build:
	go build -o ./build/discord-runner ./main.go

.PHONY: test
	go test ./...

.PHONY: docker
docker:
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output=type=registry \
		--no-cache \
		--tag shanduur/discord-runner:$(cat ./version.txt) \
		--tag shanduur/discord-runner:dev \
		.
