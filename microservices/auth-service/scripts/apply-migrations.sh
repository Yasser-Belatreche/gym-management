#!/bin/bash

set -e

# Check if atlas command exists
if ! command -v atlas &> /dev/null
then
    echo "atlas could not be found, downloading..."
    curl -sSf https://atlasgo.sh | sh
fi

atlas migrate apply --env gorm