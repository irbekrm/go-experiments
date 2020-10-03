# gRPC

## Protocol buffers

* A platform-independent format for encoding data for storage and transfer

* Define a schema in a .proto file in terms of message types (need to implement proto.Message interface) and compile with `protoc` with a language specific plugin to generate the corresponding types (with methods to set their fields etc) in that language

1. `brew install protobuf` // install the `protoc` compiler
2. `go install github.com/golang/protobuf/protoc-gen-go` // install the Go protocol buffers plugin v1 (v2 is at `google.golang.org/protobuf/cmd/protoc-gen-go`, currently not stable)

3. `protoc -I [proto_dir_path] --go_out=[output_path] [proto_file_name]` // Parse protobuf schema file & generate Go code

4. `cat [FILE] | protoc --decode_raw` // View contents of file containing protocol messages

`hexdump -c [FILE]` // Useful for debugging- see special characters in a file with data encoded using protocol buffers

Docs etc:

* https://developers.google.com/protocol-buffers

## gRPC

### With v1 (github.com/golang/protobuf/protoc-gen-go)
1. Run `protoc -I=[proto_dir_path] --go_out=plugins=grpc:[output_path] [proto_file_name]`

### With v2 (google.golang.org/protobuf/cmd/protoc-gen-go, currently not stable)
// get the go-grpc plugin to generate services
1. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

// use --go_out= to generate go code for messages and --go-grpc_out to generate go code for services
2. `protoc -I=[proto_dir_path] --go-grpc_out=[output_path --go_out=[output_path] [proto_file_name]`

### Docs etc
* https://www.grpc.io/docs/

### Testing a gRPC connection
[grpcurl](https://github.com/fullstorydev/grpcurl) is great CLI tool for testing gRPC server that supports [server reflection](https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md). Works also with Unix Domain sockets (i.e for testing CSI driver socket- `grpcurl -unix -plaintext /csi/csi.sock list`)

`grpcurl -unix -plaintext -msg-template /csi/csi.sock   describe .csi.v1.NodePublishVolumeRequest` // example- get json template of a protobuf message

https://www.youtube.com/watch?v=dDr-8kbMnaw&ab_channel=GopherAcademy&t=1s a talks about `grpcurl` and protobuf and grpc reflection, second part has actual `grpcurl` examples too

There is also [gRPC CLI](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md)