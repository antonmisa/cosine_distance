FROM golang:1.20.4

COPY . /app

WORKDIR /app

RUN go mod download

CMD ["go", "test", "-benchmem", "-bench", "."]