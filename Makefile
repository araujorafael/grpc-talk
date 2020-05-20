.PHONY: proto stress
olar: 
	echo "olar
proto:
	protoc proto/cypher.proto --proto_path=./proto --go_out=plugins=grpc:${PWD}/proto
