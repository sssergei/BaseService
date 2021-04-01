module github.com/sssergei/BaseService

go 1.16

replace github.com/sssergei/BaseService/tree/main/service => ../service

require (
	golang.org/x/net v0.0.0-20210331212208-0fccb6fa2b5c
	google.golang.org/grpc v1.36.1
)
