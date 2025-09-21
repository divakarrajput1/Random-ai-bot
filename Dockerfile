FROM golang:1.22 AS builder
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/ai-service .

FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app
COPY --from=builder /out/ai-service /app/ai-service
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/ai-service"]
