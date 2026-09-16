FROM --platform=$BUILDPLATFORM golang:1.25-bookworm AS build

WORKDIR /app

# Dependencies first — this layer is reused until go.mod/go.sum change.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# swag is a tool dependency, so the version and its checksum come from
# go.mod/go.sum rather than the command line. Runs on the build platform, so no
# GOARCH here.
RUN go tool swag init --parseDependency -q -g ./cmd/api/main.go -o ./cmd/api/docs

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -o ./build/rayls-ops-api ./cmd/api/main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/build/rayls-ops-api /app/rayls-ops-api
COPY --from=build /app/migrations /app/migrations
COPY --from=build /app/migrations-identity /app/migrations-identity
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/rayls-ops-api", "run"]
