# Builder - stage 1 of 2
FROM golang:alpine as builder
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o zapLoggerExample

# Executor - stage 2 of 2
FROM alpine:latest
WORKDIR /src/
COPY --from=builder /src/zapLoggerExample .
CMD ["./zapLoggerExample"]
