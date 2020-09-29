module github.com/tawiparkP/grpcServer

go 1.15

require (
	golang.org/x/net v0.0.0-20200923182212-328152dc79b1
	google.golang.org/grpc v1.32.0
)

require (
	github.com/tawiparkP/protofiles v0.0.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/tawiparkP/protofiles => ../protofiles
