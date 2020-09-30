# gRPC

## Protocol buffers

* A platform-independent format for encoding data for storage and transfer

* Define a schema in a .proto file in terms of message types (need to implement proto.Message interface) and compile with `protoc` with a language specific plugin to generate the corresponding types (with methods to set their fields etc) in that language

1. `brew install protobuf` // install the `protoc` compiler
2. `go install github.com/golang/protobuf/protoc-gen-go` // install the Go protocol buffers plugin v1 (v2 is at `google.golang.org/protobuf/cmd/protoc-gen-go`, currently not stable)

3. `protoc -I [proto_dir_path] --go_out=[output_path] [proto_file_name]` // Parse protobuf schema file & generate Go code

4. `cat [FILE] | protoc --decode_raw` // View contents of file containing protocol messages

`hexdump -c [FILE]` // Useful for debugging- see special characters in a file with data encoded using protocol buffers

## gRPC

### With v1 (github.com/golang/protobuf/protoc-gen-go)
1. Run `protoc -I=[proto_dir_path] --go_out=plugins=grpc:[output_path] [proto_file_name]`

### With v2 (google.golang.org/protobuf/cmd/protoc-gen-go, currently not stable)
// get the go-grpc plugin to generate services
1. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

// use --go_out= to generate go code for messages and --go-grpc_out to generate go code for services
2. `protoc -I=[proto_dir_path] --go-grpc_out=[output_path --go_out=[output_path] [proto_file_name]`

### Testing a gRPC connection
TO ADD: can be done with protoc somehow