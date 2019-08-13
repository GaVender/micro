FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

COPY consignment.json .
COPY cli .

CMD ["./cli"]