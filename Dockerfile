FROM golang:1.24-bookworm AS build

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init --parseDependency -q -g ./cmd/api/main.go -o ./cmd/api/docs

RUN CGO_ENABLED=0 go build -o ./build/rayls-ops-api ./cmd/api/main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/build/rayls-ops-api /app/rayls-ops-api
COPY --from=build /app/migrations /app/migrations
COPY --from=build /app/migrations-identity /app/migrations-identity
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/rayls-ops-api", "run"]
