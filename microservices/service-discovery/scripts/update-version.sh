#!/bin/bash

# a script that changes the project version, tag it and push it to github
# Usage: ./scripts/update-version.sh <version> <description>
# Example: ./scripts/update-version.sh 1.0.0 "first version"

set -e

# Check if the version is passed as argument
if [ -z "$1" ]
  then
    echo "Please provide a version number and a description as arguments"
    echo "Usage: ./scripts/update-version.sh <version> <description>"
    exit 1
fi

# Check if the version description is passed as argument
if [ -z "$2" ]
  then
    echo "Please provide a version description as an argument"
    echo "Usage: ./scripts/update-version.sh <version> <description>"
    exit 1
fi

# Update the version in package.json
sed -i "s/\"version\": \".*\"/\"version\": \"$1\"/g" package.json


git add .

git commit -m "$2"

git push origin master

echo "Version updated to $1"
