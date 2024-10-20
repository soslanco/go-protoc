# go-protoc
The protocol buffer compiler **protoc** with plugins: **protoc-gen-go**, **protoc-gen-go-grpc**, **protoc-gen-grpc-gateway**, **protoc-gen-openapiv2** and **.proto** files.

### Usage
#### Generate gRPC code and swagger.json file.
```sh
docker run --rm -v $(pwd):/code -w /code soslanco/go-protoc:latest \
  --go_out . --go_opt paths=source_relative \
  --go-grpc_out . --go-grpc_opt paths=source_relative \
  --grpc-gateway_out . \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --openapiv2_out . \
  --openapiv2_opt logtostderr=true \
  helloworld.proto
```
output:
```
helloworld.pb.go       
helloworld.pb.gw.go    
helloworld.swagger.json
helloworld_grpc.pb.go  
```
#### Inject custom tags
```sh
docker run --rm -v $(pwd):/code -w /code --entrypoint protoc-go-inject-tag soslanco/go-protoc:latest -input="*.pb.go"
```

#### Tips
Generate API documentation.
```sh
docker run --rm -v $(pwd):/data ghcr.io/redocly/redoc/cli build helloworld.swagger.json -o helloworld.html
```
