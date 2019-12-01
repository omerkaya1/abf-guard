# The first stage: compile integration-test binary
FROM golang:1.13-alpine as dependency-builder
WORKDIR /opt/abfg
COPY ./test/integration-test .
RUN go mod download

# The second stage: compile abfg-integration-test binary
FROM dependency-builder as app-builder
ENV APP_NAME abfg-integration-test
WORKDIR /opt/${APP_NAME}
COPY --from=dependency-builder /opt/abfg .
RUN CGO_ENABLED=0 go test -c -o ./${APP_NAME} .

# The third stage: copy the abfg-integration-test binary to another container
FROM alpine:3.9
ENV APP_NAME abfg-integration-test
LABEL name=${APP_NAME} maintainer="o.kaya" version="0.1"
WORKDIR /opt/${APP_NAME}
COPY --from=app-builder /opt/abfg-integration-test/abfg-integration-test .
COPY --from=app-builder /opt/abfg-integration-test/features ./features
