FROM golang:1.13 AS builder
WORKDIR /go/src/github.com/yumuranaoki/date
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/yumuranaoki/date/app /app
ARG APIKEY
ENV APIKEY=${APIKEY}
EXPOSE 8080
ENTRYPOINT [ "./app" ]
