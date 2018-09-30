### go-http-echo

This app responds to requests with the http request encoded in json.

#### Configuration

The following are configurable by environment variables:

`HTTP_PORT` Defaults to 8080.
`TLS_PORT`  Setting this port also enables TLS.
`TLS_CERT`  Defaults to `./server.crt`.
`TLS_KEY`   Defaults to `./server.key`.

#### Generate a TLS key/cert
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

#### Running locally

```$ go run main.go```
or
```docker-compose up```
