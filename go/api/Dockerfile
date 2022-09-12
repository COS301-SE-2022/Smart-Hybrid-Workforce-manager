# build image
FROM golang:1.18.6-alpine3.16 AS builder

WORKDIR /opt/arche-api

RUN apk add git
RUN wget -O - -q https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s v2.13.1

ADD ./api/go.mod ./api/
ADD ./api/go.sum ./api/
ADD ./lib/go.mod ./lib/
ADD ./lib/go.sum ./lib/
WORKDIR /opt/arche-api/api
RUN go mod download

WORKDIR /opt/arche-api
ADD . .

WORKDIR /opt/arche-api/api

RUN go build -o arche-api
RUN ../bin/gosec ./...

# final image
FROM alpine:3.16
RUN apk add bash
RUN adduser --disabled-password -h /opt/arche-api -G tty --shell /bin/bash arche-api
USER arche-api
WORKDIR /opt/arche-api
COPY --chown=arche-api:root --from=builder /opt/arche-api/api/arche-api ./
RUN chmod 755 /opt/arche-api
EXPOSE 8080
ENTRYPOINT ["./arche-api", "--config", "/run/secrets/config.json", "--port", "8080"]