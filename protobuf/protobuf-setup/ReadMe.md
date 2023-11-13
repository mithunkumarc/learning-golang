commands to run

1. creating project : go mod init my-protobuf
2. manual way to compile : my-protobuf> protoc --go_out=. ./proto/basic/*.proto
3. compile proto file (recommended): my-protobuf> protoc --go_opt=module=my-protobuf --go_out=. ./proto/basic/*.proto
4. pull packages : go mod tidy
5. ??? : go mod vendor


proto file structure
    
    syntax = "proto3";
    
    option go_package = "my-protobuf/protogen/basic"; // where compiled protobuf file should store
    
    message Hello {
        string name = 1;
    }
