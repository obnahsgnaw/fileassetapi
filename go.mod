module github.com/obnahsgnaw/fileassetapi

go 1.19

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230726155614-23370e0ffb3e
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

replace github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 => github.com/obnahsgnaw/grpc-gateway/v2 v2.16.0

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
)
