API User Golang Test
====================

This repo is simple CRUD User with Golang. In this case, I use *Echo* for web framework.

Before you try to test the API, we suggest you add some of these in the package repository GO

```
github.com/labstack/echo
github.com/mattn/go-sqlite3
```
-------------------------------------------------------------------------

## API Documentation

**All responses have been encoded in JSON and have the appropriate Content-Type header**

### GET /users

```
GET /users
Content-Type: "application/json"
```

##### Returns

**200** response with show all users data

```
{
    "result": [
        {
            "id": 1,
            "name": "example",
            "email": "example@mail.com",
            "handphone": "+62123456789",
            "status": 1
        },
    ]
}
```

### POST /users

```
POST /users
Content-Type: "application/json"

{
	"name" : "example",
	"email" : "example@mail.com",
	"handphone" : "+6123456789",
	"status" : 1
}
```

Attribute | Description
--------- | -----------
name      | user name
email     | user email
handphone | handphone user with ext. +62
status    | status user ex. 0 or 1

##### Returns:

**201** Created

```
{
    "message": "success created user",
    "result": {
        "id": 0,
        "name": "example",
        "email": "example@mail.com",
        "handphone": "+6123456789",
        "status": 1
    }
}
```
