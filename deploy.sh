#!/bin/bash

# Check if docker is installed
if ! [ -x "$(command -v docker)" ]; then
  echo 'Unable to find docker command, please install Docker (https://www.docker.com/) and retry' >&2
  exit 1
fi

. deviceid 

export DEVICE_ID=${DEVICEID}

# Create network by Device ID
[ ! "$(docker network ls | grep ${DEVICE_ID})" ] && docker network create -d overlay --attachable ${DEVICE_ID}
# Create network func_function if doesn't exists
[ ! "$(docker network ls | grep func_functions)" ] && docker network create -d overlay --attachable func_functions

# Deploy the docker stack with device name
docker stack deploy --compose-file docker-compose.yml ${DEVICE_ID}
