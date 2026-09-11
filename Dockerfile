# ---- Stage 1: Build ----
# Build context must be the workspace root (planx-lab/) so that local replace
# directives in go.mod can resolve sibling modules.
FROM golang:1.25-alpine AS builder

WORKDIR /build

# Copy dependency modules first for layer caching.
COPY planx-proto/ /build/planx-proto/
COPY planx-sdk-go/ /build/planx-sdk-go/

# Copy plugin module and download deps.
COPY planx-plugin-processor-passthrough/go.mod planx-plugin-processor-passthrough/go.sum ./planx-plugin-processor-passthrough/
RUN cd planx-plugin-processor-passthrough && go mod download

# Copy plugin source and build.
COPY planx-plugin-processor-passthrough/ ./planx-plugin-processor-passthrough/
RUN cd planx-plugin-processor-passthrough && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /processor-passthrough ./cmd/plugin

# ---- Stage 2: Runtime ----
FROM alpine:3.22

RUN apk add --no-cache ca-certificates

RUN addgroup -S planx && adduser -S -G planx planx

WORKDIR /plugin

COPY --from=builder /processor-passthrough /plugin/processor-passthrough

RUN chown -R planx:planx /plugin

USER planx
