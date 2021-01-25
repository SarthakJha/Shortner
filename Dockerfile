FROM golang:1.15 as build

WORKDIR /app
COPY ./ /app

RUN env GOOS=linux GOARCH=386 go build -o bin/shorty-linux .

FROM alpine as runtime

COPY --from=build /app/bin/shorty-linux /
COPY --from=build /app/dockerenv.example /
RUN cp dockerenv.example .env

CMD ["/shorty-linux"]
