FROM golang:1.23 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build .

FROM debian:bookworm-slim AS run

COPY --from=build /app/uptime /bin/uptime
EXPOSE 8080
CMD [ "uptime" ]
