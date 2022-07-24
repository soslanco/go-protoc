FROM golang:latest as build

ENV GOBIN=$GOPATH/bin

RUN apt-get update && apt-get install -yf --no-install-recommends unzip dos2unix

ENV PROTODIR=/proto
RUN mkdir -p $PROTODIR

# protoc, protoc .proto files
RUN PROTOBUF_VER=`curl -sI https://github.com/protocolbuffers/protobuf/releases/latest | \
    grep -i '^location: ' | dos2unix | grep -o '[0-9.]\+$'` && \
    PROTOC_ZIP_FILE=protoc-$PROTOBUF_VER-linux-x86_64.zip && \
    curl -OLqs https://github.com/google/protobuf/releases/download/v$PROTOBUF_VER/$PROTOC_ZIP_FILE && \
    unzip $PROTOC_ZIP_FILE bin/protoc && \
    unzip $PROTOC_ZIP_FILE include/*.proto -d /tmp && \
    mv /tmp/include/* $PROTODIR && rm -r /tmp/include && \
    rm -f $PROTOC_ZIP_FILE

# protoc-gen-grpc-gateway, protoc-gen-openapiv2, protoc-gen-openapiv2 .proto files
RUN GRPC_GATEWAY_VER=`curl -sI https://github.com/grpc-ecosystem/grpc-gateway/releases/latest | \
    grep -i '^location: ' | dos2unix | grep -o '[0-9.]\+$'` && \
    curl -sL https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v$GRPC_GATEWAY_VER/protoc-gen-grpc-gateway-v$GRPC_GATEWAY_VER-linux-x86_64 -o $GOBIN/protoc-gen-grpc-gateway && \
    chmod 0755 $GOBIN/protoc-gen-grpc-gateway && \
    curl -sL https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v$GRPC_GATEWAY_VER/protoc-gen-openapiv2-v$GRPC_GATEWAY_VER-linux-x86_64 -o $GOBIN/protoc-gen-openapiv2 && \
    chmod 0755 $GOBIN/protoc-gen-openapiv2 && \
    mkdir /tmp/grpc-gateway && \
    curl -sL https://github.com/grpc-ecosystem/grpc-gateway/archive/refs/tags/v$GRPC_GATEWAY_VER.tar.gz | tar -zxC /tmp/grpc-gateway --strip-components=1 && \
    cd /tmp/grpc-gateway && \
    find ./protoc-gen-openapiv2/ -name "*.proto" -type f -print0 | xargs -0 -I '{}' cp --parents '{}' $PROTODIR && \
    rm -fR /tmp/grpc-gateway

# protoc-gen-go, protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# googleapis .proto files
RUN mkdir /tmp/googleapis && \
    git clone --depth=1 https://github.com/googleapis/googleapis.git /tmp/googleapis && \
    cd /tmp/googleapis && \
    find . -name "*.proto" -type f -print0 | xargs -0 -I '{}' cp --parents '{}' $PROTODIR && \
    rm -fR /tmp/googleapis

RUN mkdir -p /grpc && \
    find $PROTODIR -type f > grpc.list && \
    ldd $GOBIN/* 2>/dev/null | grep -o "/\(usr/\)\?lib.*\.so\(\.[0-9]*\)\?" | sort | uniq >> grpc.list && \
    find $GOBIN -type f >> grpc.list && \
    tar -cvhf - -T grpc.list | tar -xvf - -C /grpc && \
    mkdir -p /grpc/usr/local && \
    mv /grpc/proto /grpc/usr/local/include && \
    mv /grpc$GOBIN /grpc/usr/local/bin && \
    rm -fR /grpc$GOPATH

FROM alpine:latest
COPY --from=build /grpc /
ENTRYPOINT ["protoc"]
