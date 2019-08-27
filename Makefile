.PHONY: proto stress

proto:
	protoc proto/cypher.proto --proto_path=./proto --go_out=plugins=grpc:${PWD}/proto
