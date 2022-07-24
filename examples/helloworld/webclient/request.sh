#!/bin/sh

# GET
curl http://localhost:8080/api/helloworld?name=WebClient

# POST
curl -X POST -H "Content-Type: application/json" \
  -d '{"prefix": "Hi"}' \
  http://localhost:8080/api/helloworld?name=WebClient
