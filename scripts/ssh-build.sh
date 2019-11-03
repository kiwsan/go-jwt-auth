#!/bin/bash

openssl aes-256-cbc -K $encrypted_decd9e9f07c0_key -iv $encrypted_decd9e9f07c0_iv -in travis_rsa.enc -out travis_rsa -d

chmod 600 travis_rsa
mv travis_rsa ~/.ssh/travis_rsa
