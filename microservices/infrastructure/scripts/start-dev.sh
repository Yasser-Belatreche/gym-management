#!/bin/bash

set -e

REBUILD=""
DETACH=""
SERVICE=""

# Check if --rebuild or -d flag is present
while (( "$#" )); do
  case "$1" in
    --rebuild)
      if [ -n "$2" ] && [ ${2:0:1} != "-" ]; then
        SERVICE=$2
        REBUILD="--build $SERVICE"
        shift 2
      else
        echo "Error: Argument for --rebuild is missing" >&2
        exit 1
      fi
      ;;
    -d)
      DETACH="-d"
      shift
      ;;
    *)
      echo "Error: Unsupported flag $1" >&2
      exit 1
      ;;
  esac
done

docker-compose -f infrastructure/docker/dev/docker-compose.yml up $REBUILD $DETACH