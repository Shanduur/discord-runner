#!/usr/bin/env bash

docker run \
    -d \
    --name="discord-runner" \
    -e "DOCKER_HOST='unix:///var/run/docker.sock'" \
    -e "DISCORD_TOKEN='TOKEN_HERE'" \
    -v /var/run/docker.sock:/var/run/docker.sock \
    shanduur/discord-runner:dev
