#!/bin/sh

ID=`docker create soslanco/go-protoc`
docker cp -a $ID:/usr/local/include - | gzip > proto.tar.gz
docker rm -v $ID
