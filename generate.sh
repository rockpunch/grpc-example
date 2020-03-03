#!/bin/bash
protoc protoc -I greet/greetpb/ greet/greetpb/greet.proto --go_out=plugins=grpc:greet/greetpb