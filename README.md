## Movies

## Prerequisites

1. Golang go1.20.7
2. Docker
3. Access Token Auth (You need to get a one from IMDB API)

## Overview

This is a service for gitting the latest movies from IMDB API, feed the database with the latest movies and expose it throgh a gRPC connection. This service fetches a new movie from IMDB every 3 seconds and if the latest movie is not exists in the database it will insert it into the database. 

## Important commands to compile the protocol buffers

```sh

$ export PATH="$PATH:$(go env GOPATH)/bin"
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/movies.proto

```

## Fill the ".env" file

```
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
ACCESS_TOKEN_AUTH=
```

## How to run?

```
$ docker-compose up
```