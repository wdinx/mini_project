FROM golang:1.22 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /mini-project

FROM ubuntu AS build-release-stage

WORKDIR /

#COPY --from=build-stage /app/.env /.env
COPY --from=build-stage /mini-project /mini-project

EXPOSE 1324

ENTRYPOINT ["./mini-project"]
