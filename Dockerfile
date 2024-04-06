# Builder
FROM golang:1.18.10-alpine3.17 AS builder

COPY ./ /go/src/quiz-registrator-api

WORKDIR /go/src/quiz-registrator-api
RUN go build -buildvcs=auto -o registrator-api ./cmd/registrator-api

# registrator-api

FROM alpine:3.17 as server

COPY --from=builder /go/src/quiz-registrator-api/registrator-api /bin/
COPY --from=builder /go/src/quiz-registrator-api/config.yaml /etc/

EXPOSE 8080

CMD ["/bin/registrator-api"]
