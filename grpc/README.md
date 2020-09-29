# gRPC

## Protocol buffers

* A platform-independent format for encoding data for storage and transfer

* Define a schema in a .proto file in terms of message types (need to implement proto.Message interface) and compile with `protoc` with a language specific plugin to generate the corresponding types (with methods to set their fields) in that language



1. `brew install protobuf` // install the `protoc` compiler
2. `go install google.golang.org/protobuf/cmd/protoc-gen-go` // install the Go protocol buffers plugin

3. `protoc --go_out=[output_path] [proto_file]` // Parse protobuf schema file & generate Go code

4. `cat [FILE] | protoc --decode_raw` // View contents of file containing protocol messages