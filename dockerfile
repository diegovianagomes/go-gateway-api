FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/app
RUN go build -o /main .

CMD ["/main"]

COPY .env /.env