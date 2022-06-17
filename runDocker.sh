#!/usr/bin/env bash

function log() {
    echo "$(date): ${*}"
}

imgName=${PWD##*/}
port=8080

log "Building image ${imgName}"
docker build -t "${imgName}" .

log "Starting ${imgName} in background on port ${port}"
docker run -d -p "${port}:${port}" "${imgName}"