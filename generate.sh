#protoc
# protoc -I proto/ proto/v1.recruitment.proto --go_out=plugins=grpc:v1_recruitment_grpc/

#gomock
mockgen -destination=pkg/repository/repository_mock.go -source=pkg/repository/base.go
