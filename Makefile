FILENAME=crimeandpunishment.txt
MESSAGE_DIR=${PWD}/protos/message
CLIENT_DIR=${PWD}/client

.PHONY: generate-protos
generate-protos:
	protoc -I protos protos/message.proto --go_out=plugins=grpc:protos 
	# protoc -I protos protos/message.proto --go-grpc_out=protos

.PHONY: start-server
start-server:
	go run main.go -start-server -protocol=tcp -port=9000 -filepath=${MESSAGE_DIR}/${FILENAME}

.PHONY: start-client
start-client:
	go run main.go -start-client -port=9000 -filepath=${CLIENT_DIR}/${FILENAME}

.PHONY: test
test:
	cd ${MESSAGE_DIR} && go test -v .

.PHONY: test-cpu-pprof
test-cpu-pprof:
	cd ${MESSAGE_DIR} && go test -v -cpuprofile cpu.pprof .

.PHONY: test-mem-pprof
test-mem-pprof:
	cd ${MESSAGE_DIR} && go test -v -memprofile mem.pprof .

.PHONY: test-bench
test-bench:
	cd ${MESSAGE_DIR} && go test -v -bench=.

.PHONY: see-cpu-pprof
see-cpu-pprof:
	cd ${MESSAGE_DIR} && go tool pprof -http=:8080 cpu.pprof


.PHONY: see-mem-pprof
see-mem-pprof:
	cd ${MESSAGE_DIR} && go tool pprof -http=:8081 mem.pprof
