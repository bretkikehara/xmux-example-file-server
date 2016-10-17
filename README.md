xmux example file server
================================================

Example static file server using [xmux](https://github.com/rs/xmux).

## Quickstart

Install this repository into your `$GOPATH`.

```sh
	$ go get github.com/bretkikehara/xmux-example-file-server
	$ cd $GOPATH/src/github.com/bretkikehara/xmux-example-file-server
	$ go run main.go
```

### Override port

The port can be overridden using the `PORT` environment variable.

```sh
	$ PORT=8081 go run main.go
```
