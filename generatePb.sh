#!/bin/bash

for x in pkg/auth/**/*.proto; do protoc --go_out=plugins=grpc,paths=source_relative:. $x; done
