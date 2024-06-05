#!/bin/bash

set -e

REBUILD=""
DETACH=""

# Check if --rebuild or -d flag is present
for arg in "$@"
do
    if [ "$arg" == "--rebuild" ]
    then
        REBUILD="--build app"
    elif [ "$arg" == "-d" ]
    then
        DETACH="-d"
    fi
done

docker-compose -f docker/dev/docker-compose.yml up $REBUILD $DETACH