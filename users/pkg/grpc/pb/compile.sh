#!/usr/bin/env sh

# Should be run in the project root directory

# Prerequisites:
# - brew install protobuf
# - go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# - go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc users/pkg/grpc/pb/users.proto \
--go_opt="Musers/pkg/grpc/pb/users.proto=users/pkg/grpc/pb;pb" \
--go_out=. \
--go-grpc_out=.  \
&& mv users/pkg/grpc/users.pb.go users/pkg/grpc/pb/users.pb.go
