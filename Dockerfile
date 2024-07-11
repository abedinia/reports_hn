FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o report_hn .

FROM scratch
COPY --from=builder /app/report_hn /report_hn
COPY --from=builder /app/config.yaml /config.yaml

RUN /report_hn migrate
RUN /report_hn seed

ENTRYPOINT ["/report_hn report"]
CMD ["--config", "/config.yaml"]