# Go Training application

The purpose of this application is to improve the following concepts:
- profiling
- concurency
- gRPC

### IDEA behind the application
- Create a **server** and **client** for the gRPC option (the client will request the content of the **crimeandpunishment.txt** file, which will be saved in a new location);
- The server will send the content line by line;
- The client will receive and save the file content line by line to the new location;
- The sending / saving of the lines will be done using workers in order to use concurency;
- There will be a unit test that executes the same process, with the added `cpuprofile` and `memprofile`;