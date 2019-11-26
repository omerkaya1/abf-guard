FROM golang:1.13-alpine
ENV APP_NAME abf-guard-test
ENV CGO_ENABLED 0
WORKDIR /opt/${APP_NAME}
COPY ./test/integration-test .
RUN go mod download
