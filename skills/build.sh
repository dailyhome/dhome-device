#!/bin/sh

echo "Building Skills"

export arch=$(uname -m)


# Dummy : A dummy skill that represent similar behaviour as switch
export eTAG="1.0.0"

export dockerfile="Dockerfile.armhf"

echo Building Skill s8sg/skill-switch-dummy:$eTAG
if [ "$arch" != "armv7l" ]
then
    export dockerfile="Dockerfile"
fi

docker build -t s8sg/skill-switch-dummy:$eTAG -f ./switch-dummy/$dockerfile ./switch-dummy/ --no-cache
