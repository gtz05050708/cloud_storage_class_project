#!/bin/bash
# shellcheck disable=SC2068
docker run \
    -v "$PWD":/work \
    --network=host \
    --add-host=localhost:host-gateway \
    surfstore-client $@
