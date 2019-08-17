API User Golang Test

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
            "name": "anton",
            "email": "anton@mail.com",
            "handphone": "+6281390232808",
            "status": 1
        },
        {
            "id": 2,
            "name": "anton",
            "email": "anton@mail.com",
            "handphone": "+6213892398294",
            "status": 1
        },
        {
            "id": 3,
            "name": "anton",
            "email": "anton@mail.com",
            "handphone": "+6281390232808",
            "status": 1
        },
        {
            "id": 4,
            "name": "andika",
            "email": "andika@gmail.com",
            "handphone": "+6213892398294",
            "status": 1
        }
    ]
}
```
