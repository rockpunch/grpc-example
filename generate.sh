#!/bin/bash
protoc -I sqrt/sqrtpb sqrt/sqrtpb/sqrt.proto --go_out=plugins=grpc:sqrt/sqrtpb