FROM golang:1.14.6 as builder

WORKDIR /workspace/user_srv

COPY . .

RUN go mod download && do mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /workspace/user_srv/vessel-service .

CMD ["./vessel-service"]
