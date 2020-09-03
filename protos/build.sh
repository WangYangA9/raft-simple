#!/bin/bash

protoc --proto_path="$PWD" --go_out=plugins=grpc:. "$PWD"/*.proto


