#!/bin/sh

proto_path="./api:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src"

go_import_paths="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor"
go_import_paths="${go_import_paths},Matomix/runtime/v1/primitive.proto=github.com/atomix/atomix-runtime/api/atomix/runtime/v1"
go_import_paths="${go_import_paths},Matomix/primitive/v1/primitive.proto=github.com/atomix/atomix-runtime/api/atomix/primitive/v1"
go_import_paths="${go_import_paths},Matomix/primitive/v1/descriptor.proto=github.com/atomix/atomix-runtime/api/atomix/primitive/v1"
go_import_paths="${go_import_paths},Matomix/primitive/meta/v1/object.proto=github.com/atomix/atomix-runtime/api/atomix/primitive/meta/v1"
go_import_paths="${go_import_paths},Matomix/primitive/meta/v1/timestamp.proto=github.com/atomix/atomix-runtime/api/atomix/primitive/meta/v1"

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
  --doc_out=api/atomix/primitive/v1 \
  --doc_opt=markdown,primitive.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/v1,plugins=grpc:api \
  api/atomix/primitive/v1/primitive.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/v1 \
  --doc_opt=markdown,descriptor.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/v1,plugins=grpc:api \
  api/atomix/primitive/v1/descriptor.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/meta/v1 \
  --doc_opt=markdown,object.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/meta/v1,plugins=grpc:api \
  api/atomix/primitive/meta/v1/object.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/meta/v1 \
  --doc_opt=markdown,timestamp.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/meta/v1,plugins=grpc:api \
  api/atomix/primitive/meta/v1/timestamp.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/counter/v1 \
  --doc_opt=markdown,service.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/counter/v1,plugins=grpc:api \
  api/atomix/primitive/counter/v1/service.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/election/v1 \
  --doc_opt=markdown,service.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/election/v1,plugins=grpc:api \
  api/atomix/primitive/election/v1/service.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/indexedmap/v1 \
  --doc_opt=markdown,service.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/indexedmap/v1,plugins=grpc:api \
  api/atomix/primitive/indexedmap/v1/service.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/list/v1 \
  --doc_opt=markdown,service.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/list/v1,plugins=grpc:api \
  api/atomix/primitive/list/v1/service.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/lock/v1 \
  --doc_opt=markdown,service.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/lock/v1,plugins=grpc:api \
  api/atomix/primitive/lock/v1/service.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/map/v1 \
  --doc_opt=markdown,service.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/map/v1,plugins=grpc:api \
  api/atomix/primitive/map/v1/service.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/set/v1 \
  --doc_opt=markdown,service.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/set/v1,plugins=grpc:api \
  api/atomix/primitive/set/v1/service.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/topic/v1 \
  --doc_opt=markdown,service.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/topic/v1,plugins=grpc:api \
  api/atomix/primitive/topic/v1/service.proto

protoc -I=$proto_path \
  --doc_out=api/atomix/primitive/value/v1 \
  --doc_opt=markdown,service.md \
  --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/atomix/primitive/value/v1,plugins=grpc:api \
  api/atomix/primitive/value/v1/service.proto
