FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN go build -o service_seeker .

EXPOSE 3000

CMD ["./service_seeker"]