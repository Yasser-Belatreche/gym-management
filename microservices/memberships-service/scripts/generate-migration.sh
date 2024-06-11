#!/bin/bash

set -e

# Check if the description is passed as argument
if [ -z "$1" ]
  then
    echo "Please provide a description for the migration as an arguemnt"
    echo "Usage: ./scripts/generate-migration.sh <description>"
    exit 1
fi

# Check if atlas command exists
if ! command -v atlas &> /dev/null
then
    echo "atlas could not be found, downloading..."
    curl -sSf https://atlasgo.sh | sh
fi

# Check if the --url flag is passed and it has a value
URL_FLAG=false
URL_VALUE=""

for arg in "$@"
do
    if $URL_FLAG
    then
        URL_VALUE="$arg"
        URL_FLAG=false
    fi

    if [ "$arg" = "--url" ]
    then
        URL_FLAG=true
    fi
done

EXTRA_ARGS=""

if [ -n "$URL_VALUE" ]
then
    EXTRA_ARGS="--var url=$URL_VALUE"
fi

atlas migrate diff --env gorm $EXTRA_ARGS "$1"