#!/bin/bash

# a script that trigger the deployment of the service-discovery service
# Usage: ./scripts/deploy.sh <version> <description>
# Example: ./scripts/deploy.sh 1.0.0 "first version"

set -e

# Check if the version is passed as argument
if [ -z "$1" ]
  then
    echo "Please provide a version number and a description as arguments"
    echo "Usage: ./scripts/deploy.sh <version> <description>"
    exit 1
fi

# Check if the version description is passed as argument
if [ -z "$2" ]
  then
    echo "Please provide a version description as an argument"
    echo "Usage: ./scripts/deploy.sh <version> <description>"
    exit 1
fi

# Get the current date and time
release_date=$(date +"%d-%m-%Y, %H:%M")

# Add release notes in RELEASE.md
echo "
### V $1 ($release_date)

$2
" >> RELEASE.md

# cd to the root directory
cd ../..

# Update the version in .github action
sed -i "s/VERSION: .*/VERSION: $1/g" .github/workflows/microservices-service-discovery-ci.yml

# Update the version in k8s config file
sed -i "s/gym-management-service-discovery:.*/gym-management-service-discovery:$1/g" microservices/infrastructure/k8s/service-discovery/service-discovery.yaml

git add .

git commit -m "$2"

git push origin master

git checkout deploy/microservices/service-discovery 2> /dev/null || git checkout -b deploy/microservices/service-discovery

git merge master

git push origin deploy/microservices/service-discovery

git checkout master

cd microservices/service-discovery

echo "Deployment of version $1 is triggered successfully!"
