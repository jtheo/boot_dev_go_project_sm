FROM golang:alpine as builder

WORKDIR /code
COPY . /code/

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o socialmedia

# ---

FROM alpine

USER nobody
WORKDIR /app
COPY --from=builder /code/socialmedia /app/socialmedia
EXPOSE 8080

ENTRYPOINT [ "/app/socialmedia" ]

CMD ["-addr", "0.0.0.0"]