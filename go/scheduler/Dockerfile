# build image
FROM golang:1.14-alpine3.11 AS builder

WORKDIR /opt/arche-scheduler

RUN apk add git
RUN wget -O - -q https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s v2.4.0

ADD ./scheduler/go.mod ./scheduler/
ADD ./scheduler/go.sum ./scheduler/
ADD ./lib/go.mod ./lib/
ADD ./lib/go.sum ./lib/
WORKDIR /opt/arche-scheduler/scheduler
RUN go mod download

WORKDIR /opt/arche-scheduler
ADD . .

WORKDIR /opt/arche-scheduler/scheduler

RUN go build -o arche-scheduler
RUN ../bin/gosec ./...

# final image
FROM alpine:3.11
RUN apk add bash
RUN adduser --disabled-password -h /opt/arche-scheduler -G tty --shell /bin/bash arche-scheduler
USER arche-scheduler
WORKDIR /opt/arche-scheduler
COPY --chown=arche-scheduler:root --from=builder /opt/arche-scheduler/scheduler/arche-scheduler ./
RUN chmod 755 /opt/arche-scheduler
EXPOSE 8080
ENTRYPOINT ["./arche-scheduler", "--config", "/run/secrets/config.json", "--port", "8080"]