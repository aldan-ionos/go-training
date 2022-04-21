.PHONY: generate-protos

generate-protos:
	#protoc -I protos protos/message.proto --go_out=plugins=grpc:protos/message 
	protoc -I protos protos/message.proto --go-grpc_out=protos
