#!/bin/sh

docker images | grep "soslanco/go-protoc" | awk '{print $3}' | xargs docker rmi

docker pull soslanco/go-protoc

VER_CUR=`docker run --rm soslanco/go-protoc --version | awk '{print $2}'`
docker build --no-cache -t soslanco/go-protoc .
VER=`docker run --rm soslanco/go-protoc --version | awk '{print $2}'`

echo "Removing intermediate images"
docker image prune -f --filter label=stage=builder

echo
if [ "$VER" = "$VER_CUR" ]; then
  echo "Same version: $VER"
else
  echo "New version: $VER_CUR -> $VER"
  docker tag soslanco/go-protoc soslanco/go-protoc:$VER
  docker push -a soslanco/go-protoc
fi
