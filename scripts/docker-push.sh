#!/bin/bash

echo "$DOCKER_PASSWORD" | docker login $DOCKER_CUSTOM_HOST -u "$DOCKER_USERNAME" --password-stdin

docker tag golang-identity-service:${TRAVIS_COMMIT} $DOCKER_CUSTOM_HOST/golang-identity-service:${TRAVIS_COMMIT}
docker push $DOCKER_CUSTOM_HOST/golang-identity-service:${TRAVIS_COMMIT}
docker tag golang-identity-service:${TRAVIS_COMMIT} $DOCKER_CUSTOM_HOST/golang-identity-service:latest
docker push $DOCKER_CUSTOM_HOST/golang-identity-service:latest
