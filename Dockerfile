FROM golang:1.22.2-alpine as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

FROM scratch as runtime

WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["./main"]