#!/bin/bash

# Check if docker is installed
if ! [ -x "$(command -v docker)" ]; then
  echo 'Unable to find docker command, please install Docker (https://www.docker.com/) and retry' >&2
  exit 1
fi

. configuration 

export DEVICE_ID=${DEVICEID}
export PORT=${PORT}

echo "Device ID : $DEVICE_ID, Gateway Port: $PORT"

# Create network by Device ID
echo "Creating Device Network : $DEVICE_ID"
[ ! "$(docker network ls | grep ${DEVICE_ID})" ] && docker network create -d overlay --attachable ${DEVICE_ID}
echo "Creating Function Network (func_functions) if not exist"
# Create network func_function if doesn't exists
[ ! "$(docker network ls | grep func_functions)" ] && docker network create -d overlay --attachable func_functions
echo "Deploying device stack"
# Deploy the docker stack with device name
docker stack deploy --compose-file docker-compose.yml ${DEVICE_ID}
echo "Device stack succesfully deployed"
