# syntax=docker/dockerfile:1

# Samakan dengan baris "go x.xx" di go.mod (boleh lebih baru, jangan lebih lama)
ARG GO_VERSION=1.25

# ---------- Stage 1: build ----------
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /src

# Salin manifest dependency dulu supaya layer "go mod download"
# tetap ter-cache selama go.mod / go.sum tidak berubah
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

# CGO dimatikan -> binary statis, aman dijalankan di alpine
# -trimpath dan -s -w -> binary lebih kecil dan tanpa path/debug symbol lokal
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux \
    go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

# ---------- Stage 2: runtime ----------
FROM alpine:3.21

# ca-certificates: koneksi TLS keluar | tzdata: dukungan zona waktu (env TZ)
# Folder /app/logs disiapkan milik user "app" untuk lumberjack (logs/<nama>.log)
RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app && adduser -S -G app app \
    && mkdir -p /app/logs \
    && chown -R app:app /app

WORKDIR /app
COPY --from=builder --chown=app:app /out/api /app/api

# Jalankan sebagai non-root
USER app
EXPOSE 8080

# Exec form (array) WAJIB: SIGTERM dari Docker langsung sampai ke binary Go
# sehingga graceful shutdown berjalan
ENTRYPOINT ["/app/api"]
