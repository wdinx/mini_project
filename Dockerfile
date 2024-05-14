FROM golang:1.22-alpine3.19 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /mini-project

FROM alpine:3.19 AS build-release-stage

WORKDIR /

#COPY --from=build-stage /app/.env /.env

COPY --from=build-stage /mini-project /mini-project

EXPOSE 3000

ENTRYPOINT ["./mini-project"]
