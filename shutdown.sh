#!/bin/bash

. deviceid 

docker stack rm ${DEVICEID}

docker network rm ${DEVICEID}
