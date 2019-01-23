FROM golang:1.8.3 as builder
WORKDIR /go/src/github.com/luckyluka/backend_go_rest_api
RUN go get -d -v golang.org/x/net/html
COPY back_end.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o back_end .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/luckyluka/backend_go_rest_api/back_end .
CMD ["./back_end"]