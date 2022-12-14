# syntax=docker/dockerfile:1

FROM golang:1.19.1-bullseye

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download

COPY src ./src
COPY statuscode ./statuscode
COPY .env ./
RUN go build -o /dormitory_system ./src/cmd/dormitory_system


EXPOSE 8091

CMD [ "/dormitory_system" ]