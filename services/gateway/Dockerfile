FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY services/shared/ services/shared/
COPY services/gateway/ services/gateway/
RUN cd services/gateway && go build -o /gateway ./cmd/gateway

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /gateway .
EXPOSE 8080
CMD ["./gateway"]
