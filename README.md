# Go Identity Service

## Quick start

### Login

Login using username and password to retrieve a token.

```bash
$ curl -X POST -d 'username=admin' -d 'password=password' localhost:8000/login
```

Response

```bash
{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE2NTA0NjYsInVzZXJuYW1lIjoiYWRtaW4ifQ.Ba0iLuWe3H-0Jq3mbzA10V-Z2UL4AA2G_InQjyeSeCs",

"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE3MzU5NjZ9.pNQfXAx3FTIoYCRA6tqX8YJQx-KksJGS_ISN2O1RPSQ"}
```

### Request

Request a user resource using the token in Authorization request header.

```bash
curl -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE2NTA0NjYsInVzZXJuYW1lIjoiYWRtaW4ifQ.Ba0iLuWe3H-0Jq3mbzA10V-Z2UL4AA2G_InQjyeSeCs' -i http://localhost:8000/me
```

Response

```bash
Welcome admin!
```

### Refresh Token
Request a new token using the refresh token.

```bash
curl -X POST -H 'Content-Type: application/json' -i http://localhost:8000/refresh-tokens --data '{"refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE3MzU5NjZ9.pNQfXAx3FTIoYCRA6tqX8YJQx-KksJGS_ISN2O1RPSQ"}'
```

Response

```bash
"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE2NTA3NTgsInVzZXJuYW1lIjoiYWRtaW4ifQ.OhbRLciqmXsNN70Oyw9hBzEGLakBkS72FDFMgUz4FWU","refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzE3MzYyNTh9.32fIs0D6rPsvHcPIiapz6NCQluBBpgPJP9UgwcdxgHM"}
```

## Resources
- https://echo.labstack.com/cookbook/jwt
- https://medium.com/monstar-lab-bangladesh-engineering/jwt-auth-in-go-dde432440924