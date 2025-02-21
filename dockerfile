FROM golang:1.23.5 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .

ARG SVC_NAME
ARG SVC_HOST
ARG SVC_PORT
ARG SWAGGER_HOST
ARG SWAGGER_PORT
ARG DB_HOST
ARG DB_PORT
ARG DB_USERNAME
ARG DB_PASSWORD
ARG DB_NAME
ARG SERVER_PORT

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main /app/cmd/server

FROM alpine

WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 8080

ENTRYPOINT ["/app/main"]
