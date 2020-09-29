module github.com/tawiparkP/encryptService

go 1.15

require (
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v3 v3.0.0-beta.2.0.20200911124113-3bb76868d194
	github.com/tawiparkP/encrypt/encryption v0.0.0
	github.com/tawiparkP/encrypt/utils v0.0.0
	go.etcd.io/etcd v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.27.0
)

replace github.com/tawiparkP/encrypt/encryption => ./protofiles

replace github.com/tawiparkP/encrypt/utils => ./utils

replace go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200329194405-dd816f0735f8

replace github.com/coreos/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
