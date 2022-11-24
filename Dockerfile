FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /docker-gs-ping

EXPOSE 5000

CMD [ "/docker-gs-ping" ]