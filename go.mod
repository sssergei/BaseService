module Service

go 1.16

replace github.com/sssergei/BaseService/ => ../service

require (
	github.com/sssergei/BaseService v0.0.0-20210401101202-be2fef385fde // indirect
	golang.org/x/net v0.0.0-20210331212208-0fccb6fa2b5c
	google.golang.org/grpc v1.36.1
	google.golang.org/protobuf v1.25.0
)
