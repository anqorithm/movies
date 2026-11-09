FROM golang:1.26

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /movies

EXPOSE 50051

CMD ["/movies"]