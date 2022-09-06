FROM golang:alpine AS builder

WORKDIR /app/

COPY . .

RUN go get -d ./...

RUN GOOS=linux GOARCH=amd64 go build ./cmd/rest-api/main.go

FROM scratch

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
