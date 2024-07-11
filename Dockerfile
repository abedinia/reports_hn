FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o report_hn .

FROM scratch
COPY --from=builder /app/report_hn /report_hn
COPY --from=builder /app/config.yaml /config.yaml
EXPOSE 8000
ENTRYPOINT ["/report_hn", "report"]