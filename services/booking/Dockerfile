FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY services/shared/ services/shared/
COPY services/DOCKERFILE_SVC/ services/DOCKERFILE_SVC/
RUN cd services/DOCKERFILE_SVC && go build -o /app-binary ./cmd/DOCKERFILE_SVC

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app-binary .
EXPOSE 8082
CMD ["./app-binary"]
