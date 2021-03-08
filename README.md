# gRPC Demo Application

This is a proof of concept application used to demonstrate gRPC.

The gRPC server contains `Shape` service with an RPC method `TrianglePerimeter`.

## Build and Run

Build executable and run in terminal 1.

```sh
# Linux
> go build ./server
> ./server

# Windows
> go build ./server
> server.exe
```

Call the server api in terminal 2.

```sh
# API v1 on Linux
> grpcurl -plaintext -v -d '{"sideA": 3, "sideB": 4, "sideC": 5}' localhost:50051 shape.v1.Shape/TrianglePerimeter

# API v1 on Windows
> grpcurl -plaintext -v -d "{\"sideA\": 3, \"sideB\": 4, \"sideC\": 5}" localhost:50051 shape.v1.Shape/TrianglePerimeter
```

Or build the client api in terminal 2 and pass `3 4 5` arguments as input.

```sh
# Linux
> go build ./client
> ./client 3 4 5

# Windows
> go build ./server
> server.exe 3 4 5
```

To rebuild the generated protobuf files.

```sh
protoc -I api --proto_path=api/shape/v1 --go_out=api --go-grpc_out=api --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative api/shape/v1/shape.proto
```