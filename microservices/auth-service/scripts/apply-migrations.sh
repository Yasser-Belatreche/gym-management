#!/bin/bash

set -e

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

if [ -z "$URL_VALUE" ]
then
    echo "Please provide a URL as an argument"
    echo "Usage: ./apply-migrations.sh --url <url>"
    exit 1
fi

atlas migrate apply --env gorm --var "url=$URL_VALUE"