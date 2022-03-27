#!/bin/sh

proto_path="./api:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src"

go_import_paths="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor"
go_import_paths="${go_import_paths},Matomix/runtime/v1/common.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/v1"
go_import_paths="${go_import_paths},Matomix/runtime/v1/controller.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/v1"
go_import_paths="${go_import_paths},Matomix/runtime/v1/runtime.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/v1"
go_import_paths="${go_import_paths},Matomix/runtime/primitive/v1/primitive.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/primitive/v1"
go_import_paths="${go_import_paths},Matomix/runtime/primitive/v1/descriptor.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/primitive/v1"
go_import_paths="${go_import_paths},Matomix/runtime/meta/v1/object.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/meta/v1"
go_import_paths="${go_import_paths},Matomix/runtime/meta/v1/timestamp.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/meta/v1"
go_import_paths="${go_import_paths},Matomix/runtime/counter/v1/options.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/counter/v1"
go_import_paths="${go_import_paths},Matomix/runtime/election/v1/options.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/election/v1"
go_import_paths="${go_import_paths},Matomix/runtime/indexed_map/v1/options.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/indexed_map/v1"
go_import_paths="${go_import_paths},Matomix/runtime/list/v1/options.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/list/v1"
go_import_paths="${go_import_paths},Matomix/runtime/lock/v1/options.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/lock/v1"
go_import_paths="${go_import_paths},Matomix/runtime/map/v1/options.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/map/v1"
go_import_paths="${go_import_paths},Matomix/runtime/set/v1/options.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/set/v1"
go_import_paths="${go_import_paths},Matomix/runtime/topic/v1/options.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/topic/v1"
go_import_paths="${go_import_paths},Matomix/runtime/value/v1/options.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/value/v1"

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/v1 \
  --doc_opt=markdown,common.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/v1,plugins=grpc:api \
  api/atomix/runtime/v1/common.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/v1 \
  --doc_opt=markdown,controller.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/v1,plugins=grpc:api \
  api/atomix/runtime/v1/controller.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/v1 \
  --doc_opt=markdown,runtime.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/v1,plugins=grpc:api \
  api/atomix/runtime/v1/runtime.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/primitive/v1 \
  --doc_opt=markdown,descriptor.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/primitive/v1,plugins=grpc:api \
  api/atomix/runtime/primitive/v1/descriptor.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/meta/v1 \
  --doc_opt=markdown,object.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/meta/v1,plugins=grpc:api \
  api/atomix/runtime/meta/v1/object.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/meta/v1 \
  --doc_opt=markdown,timestamp.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/meta/v1,plugins=grpc:api \
  api/atomix/runtime/meta/v1/timestamp.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/counter/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/counter/v1,plugins=grpc:api \
  api/atomix/runtime/counter/v1/primitive.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/counter/v1 \
  --doc_opt=markdown,options.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/counter/v1,plugins=grpc:api \
  api/atomix/runtime/counter/v1/options.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/election/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/election/v1,plugins=grpc:api \
  api/atomix/runtime/election/v1/primitive.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/election/v1 \
  --doc_opt=markdown,options.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/election/v1,plugins=grpc:api \
  api/atomix/runtime/election/v1/options.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/indexed_map/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/indexed_map/v1,plugins=grpc:api \
  api/atomix/runtime/indexed_map/v1/primitive.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/indexed_map/v1 \
  --doc_opt=markdown,options.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/indexed_map/v1,plugins=grpc:api \
  api/atomix/runtime/indexed_map/v1/options.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/list/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/list/v1,plugins=grpc:api \
  api/atomix/runtime/list/v1/primitive.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/list/v1 \
  --doc_opt=markdown,options.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/list/v1,plugins=grpc:api \
  api/atomix/runtime/list/v1/options.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/lock/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/lock/v1,plugins=grpc:api \
  api/atomix/runtime/lock/v1/primitive.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/lock/v1 \
  --doc_opt=markdown,options.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/lock/v1,plugins=grpc:api \
  api/atomix/runtime/lock/v1/options.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/map/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/map/v1,plugins=grpc:api \
  api/atomix/runtime/map/v1/primitive.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/map/v1 \
  --doc_opt=markdown,options.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/map/v1,plugins=grpc:api \
  api/atomix/runtime/map/v1/options.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/set/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/set/v1,plugins=grpc:api \
  api/atomix/runtime/set/v1/primitive.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/set/v1 \
  --doc_opt=markdown,options.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/set/v1,plugins=grpc:api \
  api/atomix/runtime/set/v1/options.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/topic/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/topic/v1,plugins=grpc:api \
  api/atomix/runtime/topic/v1/primitive.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/topic/v1 \
  --doc_opt=markdown,options.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/topic/v1,plugins=grpc:api \
  api/atomix/runtime/topic/v1/options.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/value/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/value/v1,plugins=grpc:api \
  api/atomix/runtime/value/v1/primitive.proto
protoc -I=$proto_path \
  --doc_out=api/atomix/runtime/value/v1 \
  --doc_opt=markdown,options.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/runtime/value/v1,plugins=grpc:api \
  api/atomix/runtime/value/v1/options.proto
