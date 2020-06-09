FROM golang AS builder

COPY . .
RUN CGO_ENABLED=0 go build -o /app

FROM alpine

COPY --from=builder /app /app

ENTRYPOINT ["/app"]
