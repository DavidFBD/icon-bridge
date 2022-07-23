#!/bin/sh

BASE_DIR=$(dirname $0)
. ${BASE_DIR}/../version.sh

DIST_DIR=build/contracts/javascore

build_image() {
    echo $BASE_DIR

    mkdir -p ${DIST_DIR}

    docker build -f ./docker/javascore/Dockerfile . --tag btp/javascore:latest
    docker create -ti --name javascore-dist -i btp/javascore

    docker cp javascore-dist:/dist/bmc.jar ${DIST_DIR}
    docker cp javascore-dist:/dist/irc2.jar ${DIST_DIR}
    docker cp javascore-dist:/dist/irc2Tradeable.jar ${DIST_DIR}
    docker cp javascore-dist:/dist/bsr.jar ${DIST_DIR}
    docker cp javascore-dist:/dist/bts.jar ${DIST_DIR}

    docker rm -f javascore-dist
}

build_image "$@"
