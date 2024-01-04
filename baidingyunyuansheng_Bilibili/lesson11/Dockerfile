FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn \
    && go mod tidy \
    && CGO_ENABLED=0 go build -o ingress-manager main.go

FROM alpine:3.15.3

WORKDIR /app

COPY --from=builder /app/ingress-manager .

CMD ["./ingress-manager"]
