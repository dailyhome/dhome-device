#!/bin/sh

export arch=$(uname -m)

export eTAG="latest"



export dockerfile="Dockerfile.armhf"

echo Building Skill s8sg/skill-dummy:$eTAG
if [ "$arch" != "armv7l" ]
then
    export dockerfile="Dockerfile"
fi

docker build -t s8sg/skill-dummy:$eTAG ./dummy -f ./dummy/$dockerfile ./dummy/ --no-cache



export dockerfile="Dockerfile.armhf"

echo Building Skill s8sg/skill-switch:$eTAG
if [ "$arch" != "armv7l" ]
then
    echo need armhf platform to build the skill switch
    exit -1
fi

docker build -t s8sg/skill-switch:$eTAG ./switch -f ./switch/$dockerfile ./switch/ --no-cache
