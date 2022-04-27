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

.PHONY: test
test:
	cd ${PWD}/protos/message && go test -v .

.PHONY: test-cpu-prof
test-cpu-prof:
	cd ${PWD}/protos/message && go test -v -cpuprofile cpu.prof .

.PHONY: test-mem-prof
test-mem-prof:
	cd ${PWD}/protos/message && go test -v -memprofile mem.prof .

.PHONY: test-bench
test-bench:
	cd ${PWD}/protos/message && go test -v -bench=.

