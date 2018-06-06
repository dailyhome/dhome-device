#!/bin/bash

# Check if docker is installed
if ! [ -x "$(command -v docker)" ]; then
  echo 'Unable to find docker command, please install Docker (https://www.docker.com/) and retry' >&2
  exit 1
fi

. deviceid 

export DEVICE_ID=${DEVICEID}

docker network create -d overlay --attachable ${DEVICE_ID}

docker stack deploy --compose-file docker-compose.yml ${DEVICE_ID}
