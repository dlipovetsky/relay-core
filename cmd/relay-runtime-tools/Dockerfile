FROM golang:1.14-alpine AS builder
ENV GO111MODULE on
ENV CGO_ENABLED 0
WORKDIR /build
COPY . .
RUN go build -a -installsuffix cgo -o /relay/runtime/tools/entrypoint ./cmd/relay-runtime-tools

FROM relaysh/core AS source

FROM gcr.io/distroless/base:debug-nonroot

COPY --from=builder /relay/runtime/tools/entrypoint /relay/runtime/tools/entrypoint
COPY --from=source /usr/local/bin/ni /relay/runtime/tools/ni

