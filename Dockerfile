# Builder
FROM golang:1.18.3 AS builder

ARG GITHUB_PATH
ARG BRANCH

WORKDIR /go/src/
RUN git clone --branch $BRANCH $GITHUB_PATH
WORKDIR /go/src/quiz-registrator-api
RUN make build

# registrator-api

FROM golang:1.18.3 as server

COPY --from=builder /go/src/quiz-registrator-api/registrator-api /bin/
COPY --from=builder /go/src/quiz-registrator-api/config.yaml /etc/

EXPOSE 8080

CMD ["/bin/registrator-api"]
