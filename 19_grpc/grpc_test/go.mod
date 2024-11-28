module grpc_test

go 1.22

require (
	github.com/golang/protobuf v1.5.3
	google.golang.org/grpc v1.56.3
)

require (
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230706204954-ccb25ca9f130 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace grpc_test/service => ../service
