FROM golang:1.18-alpine as build

RUN apk update && apk upgrade
RUN apk add --no-cache  git

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod tidy
COPY . .
RUN go build -o ./main

FROM alpine:3.16.0

WORKDIR /app

EXPOSE 8080

RUN apk update
RUN apk add --no-cache tzdata

ENV TZ=Asia/Makassar
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime

COPY --from=build /app/main /app/main

CMD ["./main"]