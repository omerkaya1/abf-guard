# The first stage: compile abf-guard
FROM golang:1.13-alpine as dependency-builder
ENV APP_NAME abf-guard
WORKDIR /opt/${APP_NAME}
COPY . .
RUN go mod download

# The second stage: 
FROM dependency-builder as app-builder
ENV APP_NAME abf-guard
WORKDIR /opt/${APP_NAME}
COPY --from=dependency-builder /opt/abf-guard .
RUN CGO_ENABLED=0 go build -o ./bin/abf-guard .

# The third stage: copy the abf-guard binary to another container
FROM scratch
LABEL name="abf-guard" maintainer="o.kaya" version="0.1"
WORKDIR /opt/abf-guard
COPY --from=app-builder /opt/abf-guard/bin/abf-guard ./bin/
COPY --from=app-builder /opt/abf-guard/configs/config-integration.json ./configs/
CMD ["./bin/abf-guard", "grpc-server", "-c", "./configs/config-integration.json"]
