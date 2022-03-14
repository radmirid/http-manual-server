# HTTP Manual Server

A manual web server that receives requests and sends responses.

## Installing

```
go get github.com/radmirid/http-manual-server
```

## Running

```
go run main.go
```

## Usage Example 1

`Input`
```
curl -v -X POST localhost:8080/users -d '{"id": 5, "name": "name"}'
```

`Output`
```
2022/03/07 12:10:34 | POST | /users | no ID
```

## Usage Example 2

`Input`
```
curl -v -X GET -H 'x-id:1' localhost:8080/users -d '{"id": 5, "name": "name"}'
```

`Output`
```
2022/03/14 12:10:56 | GET | /users | 1
```

## LICENSE

[MIT License](LICENSE)
