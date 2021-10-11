#protoc
protoc -I pkg/proto/ pkg/proto/test-stockbit.proto --go_out=plugins=grpc:./

#gomock
mockgen -destination=pkg/repository/repository_mock.go -source=pkg/repository/base.go
