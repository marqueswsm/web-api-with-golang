# Web API using golang

It is POC to create web apis using golang. It is not using a good architecture like hexagonal archicture. 
But, it contains a CRUD using PostgreSQL and Migrations.

### Used packages

- gorm.io/gorm
- github.com/gin-gonic/gin
- gorm.io/driver/postgres

### How to run

I'll assume that you have docker installed on your local machine. Therefore, you must execute the command:
```
docker-compose up --build
```

Then, you can execute the API:
```
go run main.go
```
> It is necessary to have golang installed as well

### Documentation

#### Create

```
curl --location --request POST 'http://localhost:5000/api/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Wagner"
}'
```

You will receive:
```
{
    "id": 1,
    "name": "Wagner",
    "created": "2021-08-05T10:17:50.026339526-07:00",
    "updated": "2021-08-05T10:17:50.026339526-07:00",
    "deleted": null
}
```

#### Get by ID

```
curl --location --request GET 'http://localhost:5000/api/v1/users/1'
```

You will receive:
```
{
    "id": 1,
    "name": "Wagner",
    "created": "2021-08-05T10:56:32.197227-07:00",
    "updated": "2021-08-05T10:56:32.197227-07:00",
    "deleted": null
}
```

#### Get all

```
curl --location --request GET 'http://localhost:5000/api/v1/users'
```

You will receive:
```
[
    {
        "id": 1,
        "name": "Wagner",
        "created": "2021-08-05T10:17:50.026339-07:00",
        "updated": "2021-08-05T10:17:50.026339-07:00",
        "deleted": null
    }
]
```

#### Update
```
curl --location --request PATCH 'http://localhost:5000/api/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "name": "Wagner marques"
}'
```

You will receive
```
{
    "id": 1,
    "name": "Wagner marques",
    "created": "2021-08-05T10:56:32.197227-07:00",
    "updated": "2021-08-05T10:56:32.197227-07:00",
    "deleted": null
}
```

### Delete
```
curl --location --request DELETE 'http://localhost:5000/api/v1/users/1'
```

You will receive a status 204

### How to improve

The update endpoint is receiving the ID in the request body. Try to change it and send the user id using the request path, as did in the GET by id. ;)
