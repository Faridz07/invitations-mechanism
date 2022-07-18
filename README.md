# Invitations Mechanism
Developed by Reyga Elang Fariz Virgiawan.

---

## High Level Flow
[![Services](https://i.postimg.cc/WzkH48Gf/t-drawio.png)](https://github.com/Faridz07/invitations-mechanism) 

## Features
- Admin : Basic Authentications (Register, Login), Generate & history invitations code,
- Users : Login with invitations code,
- Public : Validate invitations code.

## Tech

This Coding Test uses a number of open source projects to work properly:
- [Gin Web Framework](https://github.com/gin-gonic/gin) - Restful API for Performance and good productivity.
- [GORM V.2](https://gorm.io/gorm)  - The fantastic ORM library for Golang.
- [Viper](https://github.com/spf13/viper)  - Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats..
- [JWT-Go](github.com/golang-jwt/jwt/v4) - JWT.io has a great introduction to JSON Web Tokens, supports the parsing and verification as well as the generation and signing of JWTs.
- [Google/uuid](https://github.com/google/uuid) The uuid package generates and inspects UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.
- [BCrypt](golang.org/x/crypto/bcrypt) Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- [GoLang](https://golang.org/) - evented I/O for the backend. Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
- [Postman](https://www.postman.com/) - Postman is an API platform for building and using APIs. Postman simplifies each step of the API lifecycle and streamlines collaboration so you can create better APIs—faster.
- [Docker](https://www.docker.com/) - Create images and container for your application.
- [Postgres](https://www.postgresql.org/) - PostgreSQL is a powerful, open source object-relational database system with over 30 years of active development that has earned it a strong reputation for reliability, feature robustness, and performance.
- [Redis](https://github.com/go-redis/redis) - The open source, in-memory data store used by millions of developers as a database, cache, streaming engine, and message broker.
- [Logrus](https://github.com/sirupsen/logrus) - Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger. 


## Intallation
``` 
# cd c:/user/{user}/go/src 
# git clone https://github.com/Faridz07/invitations-mechanism.git
# cd invitations-mechanism
```

## Run
```  
# docker compose up
```

## API Documentations ([Postman Collections](https://www.getpostman.com/collections/673c407472ef383199ef))

```  
https://www.getpostman.com/collections/673c407472ef383199ef
```

## List API

**Register**

>Curl: 
```  
curl --location --request POST 'localhost:8000/api/v1/admin/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username" : "testuser",
    "email" : "testuser@gmail.com",
    "password" : "12345AA",
    "confirm_password" : "12345AA"
}'
```

>Response success
```  
{
    "status": "ok",
    "message": "success",
    "data": null
}
```

>Response failed : user already exist
```  
{
    "status": "error",
    "message": "user already exist"
}
```

**Login**

>Curl: 
```  
curl --location --request POST 'localhost:8000/api/v1/admin/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "testuser@gmail.com",
    "password": "12345AA"
}'
```

>Response success
```  
{
    "status": "ok",
    "message": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxMzk0NTQsImlzcyI6ImFwcHMxLjAiLCJpZCI6IjIyNDgwNjU3LTg2YmQtNDg1NC1hYTlkLTJjNjAzMzA3OWY0NCIsInVzZXJuYW1lIjoidGVzdHVzZXIiLCJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsInJvbGUiOiJhZG1pbiJ9.sX5RNX1sBEbJ4PpmN2o75mRSmKW8KFYf-2vPNMai5oY",
        "expiredAt": "2022-07-18T17:17:34+07:00"
    }
}
```

>Response failed : email or password doesn't match!
```  
{
    "status": "error",
    "message": "email or password doesn't match!"
}
```

**Generate Invitations Code**

>Curl: 
```  
curl --location --request GET 'localhost:8000/api/v1/admin/invitation/generate' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxMzkwNTksImlzcyI6ImFwcHMxLjAiLCJpZCI6IjIyNDgwNjU3LTg2YmQtNDg1NC1hYTlkLTJjNjAzMzA3OWY0NCIsInVzZXJuYW1lIjoidGVzdHVzZXIiLCJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsInJvbGUiOiJhZG1pbiJ9.2Mx-0lwtdw7HOLcyXzoIIXpa7-YlhFV9F1AAz6u1uP0'
```

>Response success
```  
{
    "status": "ok",
    "message": "success",
    "data": {
        "code": "H9O9krYYZCo",
        "status": "active",
        "expired_at": "2022-07-25T16:22:11.4485684+07:00"
    }
}
```

>Response failed : something when wrong, please try again!
```  
{
    "status": "error",
    "message": "something when wrong, please try again!"
}
```


**History**

>Curl: 
```  
curl --location --request GET 'localhost:8000/api/v1/admin/invitation/history?page=1&size=5' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxMzk4MDEsImlzcyI6ImFwcHMxLjAiLCJpZCI6IjIyNDgwNjU3LTg2YmQtNDg1NC1hYTlkLTJjNjAzMzA3OWY0NCIsInVzZXJuYW1lIjoidGVzdHVzZXIiLCJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsInJvbGUiOiJhZG1pbiJ9.fDprXoVHEBW2UgizTzoxwGVz2EowMaCOUZSJt3KC2fc'
```

>Response success
```  
{
    "status": "ok",
    "message": "success",
    "data": {
        "paginations": {
            "page": 1,
            "size": 5,
            "total_record": 6,
            "total_page": 2
        },
        "data": [
            {
                "id": "0e6e35f4-e779-4401-9149-f566f6a2bbe0",
                "code": "EXeCt0h",
                "status": "active",
                "created_by": "22480657-86bd-4854-aa9d-2c6033079f44",
                "created_at": "2022-07-18T16:22:27.594914+07:00",
                "expired_at": "2022-07-25T16:22:27.594914+07:00"
            },
            {
                "id": "ff8cb023-2437-4a74-9d8e-2f1250f96ffd",
                "code": "QgDFM81I781",
                "status": "active",
                "created_by": "22480657-86bd-4854-aa9d-2c6033079f44",
                "created_at": "2022-07-18T16:22:27.193052+07:00",
                "expired_at": "2022-07-25T16:22:27.193052+07:00"
            },
            {
                "id": "1823b9c1-ba5d-4e3e-8490-e221edfafae3",
                "code": "RDI00LEqN8E",
                "status": "active",
                "created_by": "22480657-86bd-4854-aa9d-2c6033079f44",
                "created_at": "2022-07-18T16:22:26.713225+07:00",
                "expired_at": "2022-07-25T16:22:26.713225+07:00"
            },
            {
                "id": "9be0ba5e-fd7f-475c-b3c3-3e6b014c4703",
                "code": "pPqA3imsJ",
                "status": "active",
                "created_by": "22480657-86bd-4854-aa9d-2c6033079f44",
                "created_at": "2022-07-18T16:22:26.322136+07:00",
                "expired_at": "2022-07-25T16:22:26.322136+07:00"
            },
            {
                "id": "8f795135-631a-4ae2-b1b2-b26cffe0f542",
                "code": "H9O9krYYZCo",
                "status": "active",
                "created_by": "22480657-86bd-4854-aa9d-2c6033079f44",
                "created_at": "2022-07-18T16:22:11.448568+07:00",
                "expired_at": "2022-07-25T16:22:11.448568+07:00"
            }
        ]
    }
}
```

>Response failed : something when wrong, please try again!
```  
{
    "status": "error",
    "message": "something when wrong, please try again!"
}
```

>Response failed : failed to get invitation, try again!
```  
{
    "status": "error",
    "message": "failed to get invitation, try again!"
}
```

**Validations invitations code**

>Curl: 
```  
curl --location --request GET 'localhost:8000/api/v1/invitation/ocnfnZ7xuqU'
```

>Response success
```  
{
    "status": "ok",
    "message": "success",
    "data": {
        "code": "ocnfnZ7xuqU",
        "status": "active",
        "expired_at": "2022-07-25T16:11:08.597441+07:00"
    }
}
```

>Response failed : invalid invitation code, try again!
```  
{
    "status": "error",
    "message": "invalid invitation code, try again!"
}
```

**Login with invitation code**

>Curl: 
```  
curl --location --request GET 'localhost:8000/api/v1/invitation/ocnfnZ7xuqU'
```

>Response success
```  
{
    "status": "ok",
    "message": "success",
    "data": "Login Successful"
}
```

>Response failed : invalid invitation code, you have 2 chances left!
```  
{
    "status": "error",
    "message": "invalid invitation code, you have 2 chances left!"
}
```

>Response failed : too many failed login attempts, please try again in 30m0s!
```  
{
    "status": "error",
    "message": "too many failed login attempts, please try again in 30m0s!"
}
```

## Sampe Log
``` 
{"context":"main","level":"info","msg":"apps running at port 8000","time":"2022-07-18T16:30:36+07:00"}
{"level":"info","msg":"","request":{"xid":"ba20b1d4-2990-4a4d-85eb-4824e9e82358","method":"POST","url":"/api/v1/admin/register","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["138"],"Content-Type":["application/json"],"Postman-Token":["dbd24a1c-6124-4da9-bd60-d46c40920ef4"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["ba20b1d4-2990-4a4d-85eb-4824e9e82358"]},"body":{"confirm_password":"12345AA","email":"testuser2@gmail.com","password":"12345AA","username":"testuser2"}},"response":"","statusCode":"","time":"2022-07-18T16:30:39+07:00"}
{"level":"info","msg":"success","request":{"xid":"ba20b1d4-2990-4a4d-85eb-4824e9e82358","method":"POST","url":"/api/v1/admin/register","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["138"],"Content-Type":["application/json"],"Postman-Token":["dbd24a1c-6124-4da9-bd60-d46c40920ef4"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["ba20b1d4-2990-4a4d-85eb-4824e9e82358"]},"body":{"confirm_password":"12345AA","email":"testuser2@gmail.com","password":"12345AA","username":"testuser2"}},"response":{"status":"ok","message":"success","data":null},"statusCode":200,"time":"2022-07-18T16:30:39+07:00"}
{"level":"info","msg":"","request":{"xid":"0b5193bd-28ee-4741-8c1f-9fc3d3b9273d","method":"POST","url":"/api/v1/admin/register","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["138"],"Content-Type":["application/json"],"Postman-Token":["51000eb4-9155-476b-9d5e-f71674f3841e"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["0b5193bd-28ee-4741-8c1f-9fc3d3b9273d"]},"body":{"confirm_password":"12345AA","email":"testuser2@gmail.com","password":"12345AA","username":"testuser2"}},"response":"","statusCode":"","time":"2022-07-18T16:30:41+07:00"}
{"context":"InsertUser-fm","details":"user already exist","level":"error","msg":"error insert data to db","time":"2022-07-18T16:30:41+07:00"}
{"level":"error","msg":"user already exist","request":{"xid":"0b5193bd-28ee-4741-8c1f-9fc3d3b9273d","method":"POST","url":"/api/v1/admin/register","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["138"],"Content-Type":["application/json"],"Postman-Token":["51000eb4-9155-476b-9d5e-f71674f3841e"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["0b5193bd-28ee-4741-8c1f-9fc3d3b9273d"]},"body":{"confirm_password":"12345AA","email":"testuser2@gmail.com","password":"12345AA","username":"testuser2"}},"response":{"status":"error","message":"user already exist"},"statusCode":400,"time":"2022-07-18T16:30:41+07:00"}
{"level":"info","msg":"","request":{"xid":"7caf63fc-c71a-473c-925d-a3e5cf7509c2","method":"POST","url":"/api/v1/admin/login","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["68"],"Content-Type":["application/json"],"Postman-Token":["ef9785b3-3df1-49a3-8db3-bf8c3d10957e"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["7caf63fc-c71a-473c-925d-a3e5cf7509c2"]},"body":{"email":"testuser2@gmail.com","password":"12345AA"}},"response":"","statusCode":"","time":"2022-07-18T16:30:48+07:00"}
{"level":"info","msg":"success","request":{"xid":"7caf63fc-c71a-473c-925d-a3e5cf7509c2","method":"POST","url":"/api/v1/admin/login","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["68"],"Content-Type":["application/json"],"Postman-Token":["ef9785b3-3df1-49a3-8db3-bf8c3d10957e"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["7caf63fc-c71a-473c-925d-a3e5cf7509c2"]},"body":{"email":"testuser2@gmail.com","password":"12345AA"}},"response":{"status":"ok","message":"success","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxNDAyNDgsImlzcyI6ImFwcHMxLjAiLCJpZCI6ImQxMTU1M2ZiLTFlMmEtNGM1NS04NWJlLTgxMTZhYjIyZmQ2OCIsInVzZXJuYW1lIjoidGVzdHVzZXIyIiwiZW1haWwiOiJ0ZXN0dXNlcjJAZ21haWwuY29tIiwicm9sZSI6ImFkbWluIn0.ytZbFn6nCfnuxV8w7wdQyFVymMO7OYrb0juxFCtgQiU","expiredAt":"2022-07-18T17:30:48+07:00"}},"statusCode":200,"time":"2022-07-18T16:30:48+07:00"}
{"level":"info","msg":"","request":{"xid":"5e14481c-6d80-43db-9326-3fcd94b4e034","method":"GET","url":"/api/v1/admin/invitation/generate","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxNDAyNDgsImlzcyI6ImFwcHMxLjAiLCJpZCI6ImQxMTU1M2ZiLTFlMmEtNGM1NS04NWJlLTgxMTZhYjIyZmQ2OCIsInVzZXJuYW1lIjoidGVzdHVzZXIyIiwiZW1haWwiOiJ0ZXN0dXNlcjJAZ21haWwuY29tIiwicm9sZSI6ImFkbWluIn0.ytZbFn6nCfnuxV8w7wdQyFVymMO7OYrb0juxFCtgQiU"],"Connection":["keep-alive"],"Postman-Token":["69857e24-7b7b-4230-8d18-c7d4b3b3bdd7"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["5e14481c-6d80-43db-9326-3fcd94b4e034"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:30:56+07:00"}
{"level":"info","msg":"success","request":{"xid":"5e14481c-6d80-43db-9326-3fcd94b4e034","method":"GET","url":"/api/v1/admin/invitation/generate","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxNDAyNDgsImlzcyI6ImFwcHMxLjAiLCJpZCI6ImQxMTU1M2ZiLTFlMmEtNGM1NS04NWJlLTgxMTZhYjIyZmQ2OCIsInVzZXJuYW1lIjoidGVzdHVzZXIyIiwiZW1haWwiOiJ0ZXN0dXNlcjJAZ21haWwuY29tIiwicm9sZSI6ImFkbWluIn0.ytZbFn6nCfnuxV8w7wdQyFVymMO7OYrb0juxFCtgQiU"],"Connection":["keep-alive"],"Postman-Token":["69857e24-7b7b-4230-8d18-c7d4b3b3bdd7"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["5e14481c-6d80-43db-9326-3fcd94b4e034"]},"body":null},"response":{"status":"ok","message":"success","data":{"code":"qBtIbnB8C6a","status":"active","expired_at":"2022-07-25T16:30:56.2284184+07:00"}},"statusCode":200,"time":"2022-07-18T16:30:56+07:00"}
{"level":"info","msg":"","request":{"xid":"ad7b0d3f-d673-45e7-ba5a-9d5db364b417","method":"GET","url":"/api/v1/invitation/qBtIbnB8C6a","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Postman-Token":["6cbf09a0-071e-4eed-a74b-b61e6ed8b109"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["ad7b0d3f-d673-45e7-ba5a-9d5db364b417"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:01+07:00"}
{"level":"info","msg":"success","request":{"xid":"ad7b0d3f-d673-45e7-ba5a-9d5db364b417","method":"GET","url":"/api/v1/invitation/qBtIbnB8C6a","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Postman-Token":["6cbf09a0-071e-4eed-a74b-b61e6ed8b109"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["ad7b0d3f-d673-45e7-ba5a-9d5db364b417"]},"body":null},"response":{"status":"ok","message":"success","data":{"code":"qBtIbnB8C6a","status":"active","expired_at":"2022-07-25T16:30:56.228418+07:00"}},"statusCode":200,"time":"2022-07-18T16:31:01+07:00"}
{"level":"info","msg":"","request":{"xid":"6aec0eba-d98d-45d1-8061-cda317d0e62a","method":"GET","url":"/api/v1/invitation/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Postman-Token":["a088d1c9-b378-474e-8d2f-48d81cb25cf5"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["6aec0eba-d98d-45d1-8061-cda317d0e62a"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:04+07:00"}
{"context":"ValidateInvitation-fm","details":"invalid invitation code, try again!","level":"error","msg":"invalid invitation code, try again!","time":"2022-07-18T16:31:04+07:00"}
{"level":"error","msg":"invalid invitation code, try again!","request":{"xid":"6aec0eba-d98d-45d1-8061-cda317d0e62a","method":"GET","url":"/api/v1/invitation/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Postman-Token":["a088d1c9-b378-474e-8d2f-48d81cb25cf5"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["6aec0eba-d98d-45d1-8061-cda317d0e62a"]},"body":null},"response":{"status":"error","message":"invalid invitation code, try again!"},"statusCode":400,"time":"2022-07-18T16:31:04+07:00"}
{"level":"info","msg":"","request":{"xid":"373076d2-1e5e-47af-af93-de81c62eb911","method":"GET","url":"/api/v1/admin/invitation/history?page=1\u0026size=5","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxMzk4MDEsImlzcyI6ImFwcHMxLjAiLCJpZCI6IjIyNDgwNjU3LTg2YmQtNDg1NC1hYTlkLTJjNjAzMzA3OWY0NCIsInVzZXJuYW1lIjoidGVzdHVzZXIiLCJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsInJvbGUiOiJhZG1pbiJ9.fDprXoVHEBW2UgizTzoxwGVz2EowMaCOUZSJt3KC2fc"],"Connection":["keep-alive"],"Postman-Token":["06874852-7e21-4556-ae69-3bad4445a3d1"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["373076d2-1e5e-47af-af93-de81c62eb911"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:06+07:00"}
{"level":"info","msg":"success","request":{"xid":"373076d2-1e5e-47af-af93-de81c62eb911","method":"GET","url":"/api/v1/admin/invitation/history?page=1\u0026size=5","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgxMzk4MDEsImlzcyI6ImFwcHMxLjAiLCJpZCI6IjIyNDgwNjU3LTg2YmQtNDg1NC1hYTlkLTJjNjAzMzA3OWY0NCIsInVzZXJuYW1lIjoidGVzdHVzZXIiLCJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsInJvbGUiOiJhZG1pbiJ9.fDprXoVHEBW2UgizTzoxwGVz2EowMaCOUZSJt3KC2fc"],"Connection":["keep-alive"],"Postman-Token":["06874852-7e21-4556-ae69-3bad4445a3d1"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["373076d2-1e5e-47af-af93-de81c62eb911"]},"body":null},"response":{"status":"ok","message":"success","data":{"paginations":{"page":1,"size":5,"total_record":6,"total_page":2},"data":[{"id":"5244be25-e336-4805-8015-95182d4ca3a3","code":"qBtIbnB8C6a","status":"active","created_by":"d11553fb-1e2a-4c55-85be-8116ab22fd68","created_at":"2022-07-18T16:30:56.228418+07:00","expired_at":"2022-07-25T16:30:56.228418+07:00"},{"id":"0e6e35f4-e779-4401-9149-f566f6a2bbe0","code":"EXeCt0h","status":"active","created_by":"22480657-86bd-4854-aa9d-2c6033079f44","created_at":"2022-07-18T16:22:27.594914+07:00","expired_at":"2022-07-25T16:22:27.594914+07:00"},{"id":"ff8cb023-2437-4a74-9d8e-2f1250f96ffd","code":"QgDFM81I781","status":"active","created_by":"22480657-86bd-4854-aa9d-2c6033079f44","created_at":"2022-07-18T16:22:27.193052+07:00","expired_at":"2022-07-25T16:22:27.193052+07:00"},{"id":"1823b9c1-ba5d-4e3e-8490-e221edfafae3","code":"RDI00LEqN8E","status":"active","created_by":"22480657-86bd-4854-aa9d-2c6033079f44","created_at":"2022-07-18T16:22:26.713225+07:00","expired_at":"2022-07-25T16:22:26.713225+07:00"},{"id":"9be0ba5e-fd7f-475c-b3c3-3e6b014c4703","code":"pPqA3imsJ","status":"active","created_by":"22480657-86bd-4854-aa9d-2c6033079f44","created_at":"2022-07-18T16:22:26.322136+07:00","expired_at":"2022-07-25T16:22:26.322136+07:00"}]}},"statusCode":200,"time":"2022-07-18T16:31:06+07:00"}
{"level":"info","msg":"","request":{"xid":"bae96321-21ba-4ece-a9a9-0a03620e9a3d","method":"POST","url":"/api/v1/user/login/ocnfnZ7xuqUxx","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxx"],"Postman-Token":["84aaefab-9037-4334-89f2-a385c387dc62"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["bae96321-21ba-4ece-a9a9-0a03620e9a3d"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:11+07:00"}
{"level":"error","msg":"too many failed login attempts, please try again in 26m2s!","request":{"xid":"bae96321-21ba-4ece-a9a9-0a03620e9a3d","method":"POST","url":"/api/v1/user/login/ocnfnZ7xuqUxx","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxx"],"Postman-Token":["84aaefab-9037-4334-89f2-a385c387dc62"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["bae96321-21ba-4ece-a9a9-0a03620e9a3d"]},"body":null},"response":{"status":"error","message":"too many failed login attempts, please try again in 26m2s!"},"statusCode":400,"time":"2022-07-18T16:31:11+07:00"}
{"level":"info","msg":"","request":{"xid":"397277ce-1648-4f4a-a95c-4c04942a16cb","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6a","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxx"],"Postman-Token":["03da8685-3998-4ec9-95cb-bff84d663be3"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["397277ce-1648-4f4a-a95c-4c04942a16cb"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:13+07:00"}
{"level":"error","msg":"too many failed login attempts, please try again in 29m58s!","request":{"xid":"397277ce-1648-4f4a-a95c-4c04942a16cb","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6a","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxx"],"Postman-Token":["03da8685-3998-4ec9-95cb-bff84d663be3"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["397277ce-1648-4f4a-a95c-4c04942a16cb"]},"body":null},"response":{"status":"error","message":"too many failed login attempts, please try again in 29m58s!"},"statusCode":400,"time":"2022-07-18T16:31:13+07:00"}
{"level":"info","msg":"","request":{"xid":"77e571cf-e2fe-4673-9e36-81ec02d7d6bd","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6a","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["f3c25215-96d0-453a-bb4b-1a647f428548"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["77e571cf-e2fe-4673-9e36-81ec02d7d6bd"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:16+07:00"}
{"level":"info","msg":"success","request":{"xid":"77e571cf-e2fe-4673-9e36-81ec02d7d6bd","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6a","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["f3c25215-96d0-453a-bb4b-1a647f428548"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["77e571cf-e2fe-4673-9e36-81ec02d7d6bd"]},"body":null},"response":{"status":"ok","message":"success","data":"Login Successful"},"statusCode":200,"time":"2022-07-18T16:31:16+07:00"}
{"level":"info","msg":"","request":{"xid":"bd521e0b-ca4a-4345-824f-61c3e54fdd77","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6a","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["7a4032e7-a06d-48a4-9ab5-4f7d382c804b"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["bd521e0b-ca4a-4345-824f-61c3e54fdd77"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:21+07:00"}
{"level":"info","msg":"success","request":{"xid":"bd521e0b-ca4a-4345-824f-61c3e54fdd77","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6a","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["7a4032e7-a06d-48a4-9ab5-4f7d382c804b"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["bd521e0b-ca4a-4345-824f-61c3e54fdd77"]},"body":null},"response":{"status":"ok","message":"success","data":"Login Successful"},"statusCode":200,"time":"2022-07-18T16:31:21+07:00"}
{"level":"info","msg":"","request":{"xid":"6bd750d7-67a3-460f-ac69-80505f19f542","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["107cb689-2935-4398-b74d-99b36badd887"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["6bd750d7-67a3-460f-ac69-80505f19f542"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:22+07:00"}
{"level":"error","msg":"invalid invitation code, you have 2 chances left!","request":{"xid":"6bd750d7-67a3-460f-ac69-80505f19f542","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["107cb689-2935-4398-b74d-99b36badd887"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["6bd750d7-67a3-460f-ac69-80505f19f542"]},"body":null},"response":{"status":"error","message":"invalid invitation code, you have 2 chances left!"},"statusCode":400,"time":"2022-07-18T16:31:22+07:00"}
{"level":"info","msg":"","request":{"xid":"4718039a-66fe-4548-8db3-5f084e9ee0b1","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["a772f6de-ce23-46f7-a624-42dc68593909"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["4718039a-66fe-4548-8db3-5f084e9ee0b1"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:24+07:00"}
{"level":"error","msg":"invalid invitation code, you have 1 chances left!","request":{"xid":"4718039a-66fe-4548-8db3-5f084e9ee0b1","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["a772f6de-ce23-46f7-a624-42dc68593909"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["4718039a-66fe-4548-8db3-5f084e9ee0b1"]},"body":null},"response":{"status":"error","message":"invalid invitation code, you have 1 chances left!"},"statusCode":400,"time":"2022-07-18T16:31:24+07:00"}
{"level":"info","msg":"","request":{"xid":"21b3ba09-10e0-4d42-a96f-58e79fecadd8","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["be146e92-4deb-4b3f-8df6-963ba5ad0568"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["21b3ba09-10e0-4d42-a96f-58e79fecadd8"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:25+07:00"}
{"level":"error","msg":"too many failed login attempts, please try again in 30m0s!","request":{"xid":"21b3ba09-10e0-4d42-a96f-58e79fecadd8","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["be146e92-4deb-4b3f-8df6-963ba5ad0568"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["21b3ba09-10e0-4d42-a96f-58e79fecadd8"]},"body":null},"response":{"status":"error","message":"too many failed login attempts, please try again in 30m0s!"},"statusCode":400,"time":"2022-07-18T16:31:25+07:00"}
{"level":"info","msg":"","request":{"xid":"41dfd7a6-0396-49e9-830a-ad909972c210","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["b3d01461-a29d-4663-994c-bc6cdde009aa"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["41dfd7a6-0396-49e9-830a-ad909972c210"]},"body":null},"response":"","statusCode":"","time":"2022-07-18T16:31:25+07:00"}
{"level":"error","msg":"too many failed login attempts, please try again in 29m59s!","request":{"xid":"41dfd7a6-0396-49e9-830a-ad909972c210","method":"POST","url":"/api/v1/user/login/qBtIbnB8C6ax","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Connection":["keep-alive"],"Content-Length":["0"],"Deviceid":["xxxxx"],"Postman-Token":["b3d01461-a29d-4663-994c-bc6cdde009aa"],"User-Agent":["PostmanRuntime/7.29.0"],"Xid":["41dfd7a6-0396-49e9-830a-ad909972c210"]},"body":null},"response":{"status":"error","message":"too many failed login attempts, please try again in 29m59s!"},"statusCode":400,"time":"2022-07-18T16:31:25+07:00"}

```

## Thank you. 