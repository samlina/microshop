FROM golang:1.14-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk --no-cache add git

WORKDIR /app/microshop-user-cli

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o microshop-user-cli

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/microshop-user-cli/microshop-user-cli .

ENTRYPOINT ["./microshop-user-cli"]