FROM rust:1.75-slim AS builder
WORKDIR /app
COPY services/crypto/Cargo.toml services/crypto/Cargo.lock* ./
RUN mkdir src && echo "fn main() {}" > src/main.rs && cargo build --release || true
COPY services/crypto/ .
RUN cargo build --release

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /target/release/crypto-server .
EXPOSE 8093
CMD ["./crypto-server"]
