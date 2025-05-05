# Dockerfile
FROM golang:1.24 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o product-service .

FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/product-service /
ENTRYPOINT ["/product-service"]
