FROM golang:latest AS builder
WORKDIR gitlab.com/sb-cloud/player-ms-api

COPY . .
RUN go build  -o /api/api cmd/api/main.go
RUN go build  -o /api/migrate cmd/migrate/main.go


FROM gcr.io/distroless/cc:latest as base

USER nonroot:nonroot


ENV SERVER_HTTP_PORT=":8080"
ENV POSTGRES_DSN="postgres://admin:password@pg:5432/db"

COPY --from=builder /api/api /
COPY --from=builder /api/migrate /

CMD ["/migrate"]
CMD ["/api"]
