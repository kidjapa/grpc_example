#!/usr/bin/env bash

protoc --proto_path=proto --go_out=plugins=grpc:proto ./proto/service.proto