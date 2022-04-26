# Go Training application

The purpose of this application is to improve the following concepts:
- profiling
- concurency
- gRPC

### IDEA behind the application
- Create a **server** and **client** for the gRPC option (the client will request the content of the **crimeandpunishment.txt** file, which will be saved in a new location);
- The server will send the content line by line;
- The client will receive and save the file content line by line to the new location;
- The sending / saving of the lines will be done using `goroutines` in order to use concurency. These `goroutines` will be syncronized using a `WaitGroup`;
- There will be a unit test that executes the same process, with the added `cpuprofile` and `memprofile`;

### Dependencies
* `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
* `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
