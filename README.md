# Movies

## Overview

This is a service for gitting the latest movies from IMDB API, feed the database with the latest movies and expose it throgh a gRPC connection. This service fetches a new movie from IMDB every 3 seconds and if the latest movie is not exists in the database it will insert it into the database. 

## Prerequisites

1. Golang go 1.20.7 or higher
    * https://go.dev/doc/install
2. Protocol Buffer Compiler
    * https://grpc.io/docs/protoc-installation/
3. Docker
    * https://docs.docker.com/engine/install/
4. Access Token Auth (You need to get a one from IMDB API)

## Notes

### Using TMDB with Postam

If you are used to postamn "like me :)" then you can find there api here:
* https://www.postman.com/devrel/workspace/tmdb-api/overview

You can just simply fork it and use it directly from postman 💥

### Important commands to compile the protocol buffers

You have to compile the protocol buffers every time you change the pf :)

```sh

$ export PATH="$PATH:$(go env GOPATH)/bin"
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/proto/movies.proto

```

## Fill the ".env" file (if you run it locally) with your own settings

```
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
ACCESS_TOKEN_AUTH=
```

## Use the following ".env" file (if you run it using the provided docker & docker-compose)

```
DB_HOST=postgres
DB_USER=movies
DB_PASSWORD=movies
DB_NAME=movies
DB_PORT=5432
ACCESS_TOKEN_AUTH=
```

## How to run?

```
$ mv .env.example .env
$ sudo chmod -R 777 postgres_data
$ docker-compose up
```