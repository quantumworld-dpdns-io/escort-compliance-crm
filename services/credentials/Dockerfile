FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY services/shared/ services/shared/
COPY services/credentials/ services/credentials/
RUN cd services/credentials && go build -o /app-binary ./cmd/credentials

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app-binary .
CMD ["./app-binary"]
