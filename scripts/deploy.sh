#!/bin/bash

rm -rf .git
eval "$(ssh-agent -s)" && ssh-add ~/.ssh/travis_rsa && ssh -o "StrictHostKeyChecking=no" ${REMOTE_USER}@${SSH_HOST} -v exit

git init
git remote add deploy "${REMOTE_USER}@${SSH_HOST}:/opt/go-auth-service" && git config user.name "Travis CI" && git config user.email "kiwsan+travis@gmail.com" && git status && git add docker-compose.yml && git commit -m "deploy"
git push --force deploy master
ssh -l ${REMOTE_USER} ${SSH_HOST} "docker-compose -f ${REMOTE_APP_DIR}/public/docker-compose.${TRAVIS_ENVIRONMENT}.yml pull && docker-compose -f ${REMOTE_APP_DIR}/public/docker-compose.${TRAVIS_ENVIRONMENT}.yml up -d"
