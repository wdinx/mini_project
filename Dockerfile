FROM golang:1.22 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /goapp

FROM ubuntu AS build-release-stage

WORKDIR /

#COPY --from=build-stage /app/.env /.env
COPY --from=build-stage /goapp /goapp

EXPOSE 1324

ENTRYPOINT ["./goapp"]
