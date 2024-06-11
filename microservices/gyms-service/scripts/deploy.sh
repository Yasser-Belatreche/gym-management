#!/bin/bash

# a script that trigger the deployment of the gyms-service service
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
sed -i "s/VERSION: .*/VERSION: $1/g" .github/workflows/microservices-gyms-service-ci.yml

# Update the version in k8s config file
sed -i "s/gym-management-gyms-service:.*/gym-management-gyms-service:$1/g" microservices/infrastructure/k8s/gyms-service/gyms-service.yaml

git add .

git commit -m "$2"

git push origin master

git checkout deploy/microservices/gyms-service 2> /dev/null || git checkout -b deploy/microservices/gyms-service

git merge master

git push origin deploy/microservices/gyms-service

git checkout master

cd microservices/gyms-service

echo "Deployment of version $1 is triggered successfully!"


