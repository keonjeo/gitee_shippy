FROM golang:1.9.0 as builder

WORKDIR /workspace/gitee_shippy/user_srv

COPY . .

RUN go get -v -x && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /workspace/gitee_shippy/user_srv/user_srv .

CMD ["./user_srv"]
