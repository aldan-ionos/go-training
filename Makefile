FILENAME=crimeandpunishment.txt

.PHONY: generate-protos
generate-protos:
	protoc -I protos protos/message.proto --go_out=plugins=grpc:protos 
	# protoc -I protos protos/message.proto --go-grpc_out=protos

.PHONY: start-server
start-server:
	go run main.go -start-server -protocol=tcp -port=9000 -filepath=${PWD}/protos/message/${FILENAME}

.PHONY: start-client
start-client:
	go run main.go -start-client -port=9000 -filepath=${PWD}/client/${FILENAME}
