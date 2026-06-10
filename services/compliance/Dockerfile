FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY services/shared/ services/shared/
COPY services/compliance/ services/compliance/
RUN cd services/compliance && go build -o /app-binary ./cmd/compliance

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app-binary .
CMD ["./app-binary"]
