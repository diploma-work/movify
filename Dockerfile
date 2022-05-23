FROM golang:1.18.2-alpine AS builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /main ./cmd

FROM alpine:3.5
COPY --from=builder main /bin/main
EXPOSE 8080
ENTRYPOINT ["/bin/main"]