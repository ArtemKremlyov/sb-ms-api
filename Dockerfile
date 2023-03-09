FROM golang:latest AS builder

WORKDIR gitlab.com/sb-cloud/player-ms-api

COPY . .
RUN go build  -o /api/api cmd/api/main.go


FROM gcr.io/distroless/cc:latest as base

USER nonroot:nonroot

ARG SERVER_HTTP_PORT
ARG GRPC_PORT
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_PORT
ARG POSTGRES_DB

ENV SERVER_HTTP_PORT=${SERVER_HTTP_PORT}
ENV GRPC_PORT=${GRPC_PORT}
ENV POSTGRES_DSN="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@pg:${POSTGRES_PORT}/${POSTGRES_DB}"

COPY --from=builder /internal/api/api /

CMD ["/api"]
