# go-json-server

A simple json API server that serve json in a specific endpoint.

each endpoint is needed to be configured using config.json

## Usages

`static` directory is, where static web page need to be stored.

`config.json`

```json
{
  "endpoints": [
    { "route": "ping", "jsonpath": "jsons/ping.json" },
    { "route": "hello", "jsonpath": "jsons/hello.json" }
  ],
  "staticpath": "static",
  "port": "8080"
}
```

In this config, there are two endpoints.
`route` is the endpoint and `jsonpath` is the json file to be served.
All json files should be in `jsons` directory.

`staticpath` is the location of static web pages.
`port` or `PORT` env use for setting server port.

Use config.json to map the API.

After configuration done, run the application by,

```shell
go build .
./go-json-server
```

Test API with web at [http://localhost:8080](http://localhost:8080).

## Endpoints

| Methods | Endpoints  | Description                          |
| ------- | ---------- | ------------------------------------ |
| GET     | /ping      | Use for ping                         |
| GET     | /api/ping  | example endpoint                     |
| GET     | /api/hello | example endpoint                     |
| GET     | /api/\*    | create \* endpoint from config files |

## Development

```shell
cd go-json-server
go mod download
go run main.go
```

## Author

- Injamul Mohammad Mollah <mrinjamul@gmail.com>
