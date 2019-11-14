FROM golang:1.13.1-alpine3.10 AS build
WORKDIR /usr/src/SessionServer
COPY . .
RUN go build

FROM alpine:3.10 AS production
WORKDIR /app
COPY --from=build /usr/src/SessionServer/SessionServer .

EXPOSE 11000/tcp
ENTRYPOINT ["./SessionServer"]