gen:
	protoc -I protos/ protos/base_msg.proto --go_out=./protos
	protoc -I protos/ protos/a.proto --go-grpc_out=./protos --go_out=./protos
	protoc -I protos/ protos/b.proto --go-grpc_out=./protos --go_out=./protos

	cp -rf ./protos/github.com/vanh01/grpc-base/protos/* ./protos
	rm -rf ./protos/github.com
.PHONY: gen