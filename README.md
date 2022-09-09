# Go HTTP Server Template

## usage

### run

```shell
$ docker-compose up -d
$ go run ./cmd/http-server/main.go
```

### register

```shell
$ curl -XPOST \
  -H "content-type: application/json" \
  -d '{"first_name": "yu", "last_name": "yu", "email": "ultraganbayu10@gmail.com", "password": "passw0rd"}' \
  localhost:5000/register
```

### login

```shell
$ curl -XPOST \
  -H "content-type: application/json" \
  -d '{"email": "ultraganbayu10@gmail.com", "password": "passw0rd"}' \
  --dump-header - \
  localhost:5000/login
```

`Set-Cookie: ` の後をコピーし、変数に代入

```shell
$ COOKIE="..."
```

### logout

```shell
$ curl -XPOST \
  -H "content-type: application/json" \
  -H "cookie: ${COOKIE}" \
  --dump-header - \
  localhost:5000/logout
```