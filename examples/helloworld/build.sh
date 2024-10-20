#!/bin/sh

cd helloworld

# Generate gRPC code and swagger.json file
# (helloworld_grpc.pb.go, helloworld.pb.go, helloworld.pb.gw.go, helloworld.swagger.json)
docker run --rm -v $(pwd):/code -w /code soslanco/go-protoc \
  --go_out . --go_opt paths=source_relative\
  --go-grpc_out . --go-grpc_opt paths=source_relative \
  --grpc-gateway_out . \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --openapiv2_out . \
  --openapiv2_opt logtostderr=true \
  helloworld.proto

# Generate documentation
# (helloworld.html)
docker run --rm -v $(pwd):/data ghcr.io/redocly/redoc/cli build helloworld.swagger.json -o helloworld.html

# Inject tags
docker run --rm -v $(pwd):/code -w /code --entrypoint protoc-go-inject-tag soslanco/go-protoc:latest -input="*.pb.go"

cd ../client
go build

cd ../server
go build
