FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY services/shared/ services/shared/
COPY services/booking/ services/booking/
RUN cd services/booking && go build -o /app-binary ./cmd/booking

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app-binary .
CMD ["./app-binary"]
