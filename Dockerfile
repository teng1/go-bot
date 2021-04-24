FROM golang:1.16 as builder

WORKDIR /go/src/github.com/teng1/go-bot

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go install -v \
    github.com/teng1/go-bot

FROM alpine:3.13.2 AS final

RUN apk add --update --no-cache ca-certificates

RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -S app -G app

COPY --from=builder /go/bin/go-bot /usr/local/bin/go-bot

USER 1000

COPY entrypoint.sh .

EXPOSE 8080

ENTRYPOINT [ "./entrypoint.sh" ]