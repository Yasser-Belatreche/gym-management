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

atlas migrate diff --env gorm "$1"