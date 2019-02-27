FROM golang:1.12-alpine as builder

# Git is needed for Go to pull dependencies
RUN apk add --no-cache git

WORKDIR /src

# Pull dependencies into their own layers for more efficient caching.
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build -o /app/orb-badge-server ./...


FROM alpine:3.9

LABEL maintainer="ricardo@feliciano.tech"

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/orb-badge-server /app/

WORKDIR /app

CMD /app/orb-badge-server

EXPOSE 1107
