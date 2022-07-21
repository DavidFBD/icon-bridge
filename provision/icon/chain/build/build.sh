#!/bin/bash

DOCKER_REPO=${DOCKER_REPO:-localhost:5000}
DOCKER_IMG=${DOCKER_IMG:-icon-chain}
DOCKER_TAG=${DOCKER_TAG:-latest}
DOCKER_IMGTAG=${DOCKER_IMGTAG:-"$DOCKER_REPO/$DOCKER_IMG:$DOCKER_TAG"}

docker build --build-arg CONFIG_JSON="$(cat icon.config.json)" -t $DOCKER_IMGTAG .
