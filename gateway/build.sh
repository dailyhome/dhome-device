#!/bin/sh

export dockerfile="Dockerfile"
export arch=$(uname -m)

export eTAG="1.0.0"

if [ "$arch" = "armv7l" ] ; then
   dockerfile="Dockerfile.armhf"
   eTAG="1.0.0-armhf"
fi

echo Building s8sg/device-gateway:$eTAG

docker build -t s8sg/device-gateway:$eTAG . -f $dockerfile --no-cache
