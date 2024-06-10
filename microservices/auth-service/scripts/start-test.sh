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

docker-compose -f docker/test/docker-compose.yml up $REBUILD $DETACH && scripts/apply-migrations.sh --url "postgres://postgres:postgres@localhost:5432/postgres"