# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS builder
# FROM golang AS builder

WORKDIR /build
RUN adduser -u 10001 -D app-runner

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# COPY *.go ./
COPY . .

# RUN go build -o /pricefeed
# EXPOSE 8080
# CMD [ "/pricefeed" ]

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o pricefeed .
# RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -a -o pricefeed .


FROM alpine:3.10 AS final

WORKDIR /app
COPY --from=builder /build/pricefeed /app/
# move config file 
# COPY --from=builder /build/* /app/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER app-runner
ENTRYPOINT ["/app/pricefeed"]

# docker network inspect bridge