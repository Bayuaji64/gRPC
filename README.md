# gRPC
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2



protoc -Igreet/proto --go_out=. --go-grpc_out=. greet/proto/dummy.proto

protoc -Igreet/proto --go_out=. --go_opt=module=github.com/Bayuaji64/grpc-go --go-grpc_out=. --go-grpc_opt=module=github.com/Bayuaji64/grpc-go greet/proto/dummy.proto


make greet
go mod tidy


./bin/greet/server