## Simple HTTP-Server

### Start server (with default port)

```sh
go run ./src/http_server.go
# -> http://0.0.0.0:8080
```

### Start server (with specific port)

```sh
go run ./src/http_server.go -port=1234
# -> http://0.0.0.0:1234
```

## Simple REST-API

```sh
go run ./src/rest_api.go
```

### Reverse string (GET)

```sh
curl --location --request GET 'http://0.0.0.0:5678/reverse/nyanko'
```

### Create user (POST)

```sh
curl --location --request POST 'http://0.0.0.0:5678/signup' \
--form 'name="jpnykw"' \
--form 'age="19"' \
--form 'email="jpnykw@gmail.com"'
```

### Get user (GET)

```sh
curl --location --request GET 'http://0.0.0.0:5678/users/get/jpnykw'
```

## Database connection (MySQL)

First, you need to create the `.env` file and set the necessary values in the environment variables.

```sh
go run ./src/connect_db.go
```
