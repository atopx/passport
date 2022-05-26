#!/bin/sh

protoc --proto_path=./protocol --go_out=plugins=grpc:. ./protocol/*.proto
