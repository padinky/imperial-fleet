FROM golang:1.21-alpine AS builder

ENV SERVICE_PATH=./main.go

RUN apk add --update git ca-certificates tzdata

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o ./executable ${SERVICE_PATH}

# ===============================================================================================
# Start fresh from a smaller image
FROM alpine:3.16

RUN apk --no-cache add ca-certificates tzdata

RUN mkdir -p /build
RUN mkdir -p /build/log
WORKDIR /build

COPY --from=builder /app/executable /build

CMD ./executable
