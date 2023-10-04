# Drips

Nothing like a little sweat.

## Server

Start the local gRPC server with CRUD operations.

```bash
go run ./cmd/server
```

## Client

Run a cli client

```bash
go run ./cmd/client exercise-class list
```

## Development

Compile Golang language bindings.

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/drips.proto
```