proto:
	protoc --go_out=. --go-grpc_out=. pkg/pb/auth.proto

path:
	export PATH="$PATH:$(go env GOPATH)/bin"