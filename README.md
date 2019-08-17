API User Golang Test
====================

This repo is simple CRUD User with Golang. In this case, I use ***Echo*** for web framework.

Before you try to test the API, we suggest you add some of these in the package repository GO

```
github.com/labstack/echo
github.com/mattn/go-sqlite3
```

Structure in Golang:
```
- go
  |- bin
  |- pkg
  |- src
    |- api-test
    |- github.com
        |- labstack
            |- echo
        |- mattn
            |- go-sqlite3
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

**200** OK (Show All Users Data)

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

**201** Created (Success Created User)

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

##### Errors:

Error | Description
----- | ------------
400   | failed insert data
404   | ```url``` not found 


### PATCH /users/update-handphone/:id

```
PATCH /users/update-handphone/:id
Content-Type: "application/json"

{
	"handphone" : "+6123456789",
}
```

Attribute | Description
--------- | -----------
**id**    | id user
handphone | handphone user with ext. +62

##### Returns:

**200** OK (Success Update Handphone User)

```
{
    "message": "success update handphone user",
    "result": "+62123456789"
}
```

##### Errors:

Error | Description
----- | ------------
400   | failed update handphone user
404   | ```url``` not found 


### PATCH /users/update-status/:id

```
PATCH /users/update-status/:id
Content-Type: "application/json"

{
	"status" : 1,
}
```

Attribute | Description
--------- | -----------
**id**    | id user
status    | status user ex. 0 or 1 (integer)

##### Returns:

**200** OK (Success Update Status User)

```
{
    "message": "success update status user",
    "result": 1
}
```

##### Errors:

Error | Description
----- | ------------
400   | failed update status user
404   | ```url``` not found 


### DELETE /users/delete/:id

```
DELETE /users/delete/:id
Content-Type: "application/json"
```

Attribute | Description
--------- | -----------
**id**    | id user

##### Returns:

**200** OK (Success Delete User)

```
{
    "message": "success deleted user"
}
```

##### Errors:

Error | Description
----- | ------------
400   | failed deleted user
404   | ```url``` not found 
