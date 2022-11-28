FROM golang:1.19

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE 5000

CMD ["go", "run", "main.go"]