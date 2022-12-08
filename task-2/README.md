# Server/Client Api

This is a small project with:

* simple REST api server, which returns json with the list of objects
* simple client, which fetches the json from server and parses it out  

## How to run

### Server

``` bash
cd server
go get -u && go mod tidy
go run main.go
```

process is attached to the tty while running

### Client

``` bash
cd client
go get -u && go mod tidy
go run main.go
```
