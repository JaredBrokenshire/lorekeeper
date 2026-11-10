FROM golang:1.26-alpine

WORKDIR /app

# git is required by go modules; gcc + musl-dev for any CGO dependencies
RUN apk add --no-cache git gcc musl-dev

# docker-compose-wait blocks startup until the DB healthcheck passes
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

# Install dev tooling
RUN go install github.com/githubnemo/CompileDaemon@latest && \
    go install github.com/swaggo/swag/cmd/swag@v1.16.6

# Pre-download dependencies (cached unless go.mod/go.sum change)
COPY go.mod go.sum ./
RUN go mod download

# Source is mounted via the volume in docker-compose so changes are picked up
# without rebuilding the image. The COPY here ensures the image works standalone.
COPY . .

# Generate Swagger docs then start the server with live reload.
# Binary is written to /app/bin/ which is outside the mounted source tree.
CMD swag init -g cmd/api/api.go --output docs && \
    /wait && \
    CompileDaemon \
      --build="go build -o /app/bin/server ./cmd/api" \
      --command="/app/bin/server" \
      --color \
      --log-prefix=false
