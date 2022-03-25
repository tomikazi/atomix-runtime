#!/bin/sh

proto_path="./api/proto:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src"

rm -r api/docs
mkdir -p api/docs

mkdir -p api/docs/atomix/primitive/v1
mkdir -p api/docs/atomix/primitive/counter/v1
mkdir -p api/docs/atomix/primitive/lock/v1
mkdir -p api/docs/atomix/primitive/map/v1
mkdir -p api/docs/atomix/primitive/set/v1
mkdir -p api/docs/atomix/primitive/topic/v1

protoc -I=$proto_path --doc_out=api/docs/atomix/primitive/v1         --doc_opt=markdown,primitive.md api/proto/atomix/primitive/v1/primitive.proto
protoc -I=$proto_path --doc_out=api/docs/atomix/primitive/counter/v1 --doc_opt=markdown,service.md   api/proto/atomix/primitive/counter/v1/service.proto
protoc -I=$proto_path --doc_out=api/docs/atomix/primitive/lock/v1    --doc_opt=markdown,service.md   api/proto/atomix/primitive/lock/v1/service.proto
protoc -I=$proto_path --doc_out=api/docs/atomix/primitive/map/v1     --doc_opt=markdown,service.md   api/proto/atomix/primitive/map/v1/service.proto
protoc -I=$proto_path --doc_out=api/docs/atomix/primitive/set/v1     --doc_opt=markdown,service.md   api/proto/atomix/primitive/set/v1/service.proto
protoc -I=$proto_path --doc_out=api/docs/atomix/primitive/topic/v1   --doc_opt=markdown,service.md   api/proto/atomix/primitive/topic/v1/service.proto

mkdir -p api/docs/atomix/runtime/v1

protoc -I=$proto_path --doc_out=api/docs/atomix/runtime/v1 --doc_opt=markdown,controller.md api/proto/atomix/runtime/v1/controller.proto
protoc -I=$proto_path --doc_out=api/docs/atomix/runtime/v1 --doc_opt=markdown,primitive.md  api/proto/atomix/runtime/v1/primitive.proto
protoc -I=$proto_path --doc_out=api/docs/atomix/runtime/v1 --doc_opt=markdown,agent.md      api/proto/atomix/runtime/v1/agent.proto

go_import_paths="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor"
go_import_paths="${go_import_paths},Matomix/primitive/v1/primitive.proto=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/v1"
go_import_paths="${go_import_paths},Matomix/primitive/v1/descriptor.proto=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/v1"
go_import_paths="${go_import_paths},Matomix/primitive/meta/v1/object.proto=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/meta/v1"
go_import_paths="${go_import_paths},Matomix/primitive/meta/v1/timestamp.proto=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/meta/v1"

protoc -I=$proto_path --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/proto/atomix/runtime/v1,plugins=grpc:api/proto           api/proto/atomix/runtime/v1/*.proto
protoc -I=$proto_path --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/v1,plugins=grpc:api/proto         api/proto/atomix/primitive/v1/*.proto
protoc -I=$proto_path --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/meta/v1,plugins=grpc:api/proto    api/proto/atomix/primitive/meta/v1/*.proto
protoc -I=$proto_path --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/counter/v1,plugins=grpc:api/proto api/proto/atomix/primitive/counter/v1/*.proto
protoc -I=$proto_path --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/lock/v1,plugins=grpc:api/proto    api/proto/atomix/primitive/lock/v1/*.proto
protoc -I=$proto_path --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/map/v1,plugins=grpc:api/proto     api/proto/atomix/primitive/map/v1/*.proto
protoc -I=$proto_path --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/set/v1,plugins=grpc:api/proto     api/proto/atomix/primitive/set/v1/*.proto
protoc -I=$proto_path --gogofaster_out=$go_import_paths,import_path=github.com/atomix/atomix-runtime/api/proto/atomix/primitive/topic/v1,plugins=grpc:api/proto   api/proto/atomix/primitive/topic/v1/*.proto
