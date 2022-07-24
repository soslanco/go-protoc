# goprotoc
The protocol buffer compiler **protoc** with plugins: **protoc-gen-go**, **protoc-gen-go-grpc**, **protoc-gen-grpc-gateway**, **protoc-gen-openapiv2** and **.proto** files.

### Usage
```sh
docker run --rm -v $(pwd):/code -w /code soslanco/goprotoc \
  --go_out . --go_opt paths=source_relative\
  --go-grpc_out . --go-grpc_opt paths=source_relative \
  --grpc-gateway_out . \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --openapiv2_out . \
  --openapiv2_opt logtostderr=true \
  helloworld.proto
```
Output:
```
helloworld.pb.go       
helloworld.pb.gw.go    
helloworld.swagger.json
helloworld_grpc.pb.go  
```

#### Tips
Generate API documentation
```sh
docker run --rm -v $(pwd):/data ghcr.io/redocly/redoc/cli:v2.0.0-rc.72 build helloworld.swagger.json -o helloworld.html
```
