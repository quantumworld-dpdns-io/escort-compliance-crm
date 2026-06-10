FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY services/shared/ services/shared/
COPY services/auth/ services/auth/
RUN cd services/auth && go build -o /auth ./cmd/auth

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /auth .
EXPOSE 8081
CMD ["./auth"]
