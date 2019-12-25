# Go Identity Service

[![Build Status](https://travis-ci.org/kiwsan/go-jwt-auth.svg?branch=master)](https://travis-ci.org/kiwsan/go-jwt-auth)
[![CircleCI](https://circleci.com/gh/kiwsan/go-jwt-auth/tree/master.svg?style=svg)](https://circleci.com/gh/kiwsan/go-jwt-auth/tree/master)

[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=kiwsan_go-jwt-auth)](https://sonarcloud.io/dashboard?id=kiwsan_go-jwt-auth)

## Quick start

### Register
Request a new user to login.

```bash
$ curl -X POST -H 'Content-Type: application/json' -i http://localhost:8000/sign-up --data '{ "email": "admin@kiwsan.com", "password": "password", "confirm_password": "password" }'
```

Response
```bash
The user has been created.
```

### Login

Login using email and password to retrieve a token.

```bash
$ curl -X POST -d 'email=admin@kiwsan.com' -d 'password=password' localhost:8000/sign-in
```

Response

```bash
{
  "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE4MzQyMzIsInVzZXJuYW1lIjoiYWRtaW4ifQ.J8KCW98u2JMC1kqd2xStp10WTYYb9lksdR4QYtXQffc",
  "expire_date":"1571834232",
  "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE5MTk3MzJ9.6pPqklkbIrSrgT82wqS_Dn4UFs_CUk_MkSt1BdXeNvQ"
}
```

### Request user claims information

Request a user resource using the token in Authorization request header.

```bash
$ curl -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE4MzQyMzIsInVzZXJuYW1lIjoiYWRtaW4ifQ.J8KCW98u2JMC1kqd2xStp10WTYYb9lksdR4QYtXQffc' -i http://localhost:8000/me
```

Response

```bash
Welcome admin!
```

### Refresh Token
Request a new token using the refresh token.

```bash
$ curl -X POST -H 'Content-Type: application/json' -i http://localhost:8000/access-tokens/:refreshToken/refresh --data '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE5MTk3MzJ9.6pPqklkbIrSrgT82wqS_Dn4UFs_CUk_MkSt1BdXeNvQ"}'
```

Response

```bash
{
  "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE4MzQzOTgsInVzZXJuYW1lIjoiYWRtaW4ifQ.IueTxg55g0R2DG9z_I6y3ea1YCSr8pm0SO_A-9LV_vQ",
  "expire_date":"1571834398",
  "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE5MTk4OTh9.Z3z9Lz8C3nh5BbbxAMwvQYRW9wmcsgvrgFlYTrTS3og"
}
```

### Revoke Refresh Token
Request a new token using the refresh token.

```bash
$ curl -X POST -H 'Content-Type: application/json' -i http://localhost:8000/refresh-tokens/:refreshToken/revoke --data '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE5MTk3MzJ9.6pPqklkbIrSrgT82wqS_Dn4UFs_CUk_MkSt1BdXeNvQ"}'
```

Response

```bash
{
  "access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE4MzQzOTgsInVzZXJuYW1lIjoiYWRtaW4ifQ.IueTxg55g0R2DG9z_I6y3ea1YCSr8pm0SO_A-9LV_vQ",
  "expire_date":"1571834398",
  "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE5MTk4OTh9.Z3z9Lz8C3nh5BbbxAMwvQYRW9wmcsgvrgFlYTrTS3og"
}
```

## Travis CLI

Install the Travis CLI

```bash
$ gem install travis
```

```bash
$ travis login
```

Create the encrypted key and public key in your local repo:

```bash
$ ssh-keygen -t rsa -N "" -C "kiwsanthia@gmail.com" -f travis_rsa
```

```bash
$ travis encrypt-file travis_rsa --add
```

Setup Travis on the droplet

```bash
$ sudo adduser --disabled-password --gecos "" travis
```

```bash
$ sudo chown -R travis:travis /opt/www
```

```bash
$ sudo su travis
```

```bash
$ mkdir ~/.ssh
```

```bash
$ chmod 700 .ssh
```

```bash
$ vim .ssh/authorized_keys # paste content from travis_rsa.pub
```

```bash
$ chmod 600 .ssh/authorized_keys && exit
```

Prepare remote repository on the droplet
 
```bash
$ sudo su travis
```

```bash
$ cd /opt/www
```

```bash
$ mkdir .git
```

```bash
$ cd .git
```

```bash
$ git init --bare
```

```bash
$ cd hooks
```

```bash
$ vim post-receive 
```

post-receive hook

```bash
#!/bin/sh
git --work-tree=/opt/www/public/ --git-dir=/opt/www/.git checkout -f
```

```bash
$ chmod +x post-receive
```

## Resources
- https://echo.labstack.com/cookbook/jwt
- https://medium.com/monstar-lab-bangladesh-engineering/jwt-auth-in-go-dde432440924
- https://www.baeldung.com/rest-api-error-handling-best-practices
- https://docs.travis-ci.com/user/encrypting-files/
