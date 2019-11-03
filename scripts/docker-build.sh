#!/bin/bash

docker build -t golang-identity-service:${TRAVIS_COMMIT} .
docker run -d -p 127.0.0.1:8000:8000 golang-identity-service:${TRAVIS_COMMIT}
