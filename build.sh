#!/bin/sh

docker pull soslanco/go-protoc
VER_CUR=`docker run soslanco/go-protoc --version | awk '{print $2}'`
docker build --no-cache -t soslanco/go-protoc .
VER=`docker run soslanco/go-protoc --version | awk '{print $2}'`

echo
if [ "$VER" = "$VER_CUR" ]; then
  echo "Same version: $VER"
else
  echo "New version: $VER_CUR -> $VER"
  docker tag soslanco/go-protoc soslanco/go-protoc:$VER
  docker push -a soslanco/go-protoc
fi
