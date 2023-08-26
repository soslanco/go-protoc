#!/bin/sh

ID=`docker create soslanco/go-protoc`
docker cp -a $ID:/usr/local/include - | gzip > proto-include.tar.gz
docker cp -a $ID:/usr/local/bin - | gzip > proto-bin.tar.gz
docker rm -v $ID
