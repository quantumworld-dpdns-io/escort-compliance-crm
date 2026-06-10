FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY services/shared/ services/shared/
COPY services/companion/ services/companion/
RUN cd services/companion && go build -o /app-binary ./cmd/companion

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app-binary .
CMD ["./app-binary"]
