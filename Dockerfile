FROM golang:1.15-alpine as builder
WORKDIR /app
COPY . .
RUN go install -v ./...

FROM alpine:3.12
WORKDIR /app
COPY --from=builder /go/bin/parsimonious-users-api .

ENTRYPOINT ["/app/parsimonious-users-api"]
