FROM golang:1.25.1-alpine3.22 AS base
WORKDIR /go/src/app
RUN apk update
RUN apk --no-cache add curl
RUN apk --no-cache add ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -a -ldflags "-s -w" -trimpath -o main .

EXPOSE 50052

FROM base AS local
RUN go install github.com/air-verse/air@v1.61.5

CMD ["air"]

FROM alpine:3.21 AS production

RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

WORKDIR /root/

COPY --from=base /go/src/app/main .
COPY --from=base /go/src/app/.env .

CMD ["/root/main"]
