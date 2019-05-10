#!/bin/bash

for x in pkg/auth/local/*.proto; do protoc --go_out=plugins=grpc,paths=source_relative:. $x; done

for x in pkg/storage/bolt/*.proto; do protoc --go_out=paths=source_relative:. $x; done
