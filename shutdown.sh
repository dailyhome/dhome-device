#!/bin/bash

. configuration 

echo "Removing stack"
docker stack rm ${DEVICEID}
echo "Removing device network"
docker network rm ${DEVICEID}
echo "Device ${DEVICEID} teardowned successfully"
