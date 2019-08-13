FROM golang as builder

ARG PROJECT_PATH=github.com/GaVender/micro

RUN mkdir -p /go/src/$PROJECT_PATH
WORKDIR /go/src/$PROJECT_PATH

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /vessel-service service/vessel/main.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /vessel-service .

CMD ["./vessel-service"]