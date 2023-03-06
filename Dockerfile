FROM golang:latest AS builder

WORKDIR gitlab.com/sb-cloud/player-ms-api

COPY . .
RUN go build  -o /api/api cmd/api/main.go
RUN go build  -o /api/migrate cmd/migrate/main.go


FROM gcr.io/distroless/cc:latest as base

USER nonroot:nonroot

ARG SERVER_HTTP_PORT
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_PORT
ARG POSTGRES_DB

ENV SERVER_HTTP_PORT=${SERVER_HTTP_PORT}
ENV POSTGRES_DSN="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@pg:${POSTGRES_PORT}/${POSTGRES_DB}"

COPY --from=builder /api/api /
COPY --from=builder /api/migrate /

CMD ["/migrate"]
CMD ["/api"]
