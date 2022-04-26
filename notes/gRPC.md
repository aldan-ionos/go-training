# gRPC Notes

## What is gRPC?

**gRPC** is a modern open source high performance **Remote Procedure Call (RPC)** framework that can **run in any environment**. It can efficiently **connect services** in and across data centers with **pluggable support** for *load balancing, tracing, health checking and authentication*. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

---
## How does gRPC compare to REST API?

### Similarities:
* Both are communication protocols between services;
* Both use an `http` protocol; 

### Differences:
| **REST API** | **gRPC** |
| --- | --- |
| uses the `http 1` protocol | uses the `http 2` protocol |
| uses `json` messaging format | uses `protocol buffer (protobuf)` messaging format |

---
## What is protobuf?
**Protocol Buffers (Protobuf)** is a free and open-source cross-platform data **format used to serialize structured data**. It is useful in **developing programs to communicate** with each other over a network or for storing data. The method involves an interface description language that describes the structure of some data and a program that generates source code from that description for **generating** or parsing a **stream of bytes** that represents the structured data. 

---
## Why use protobuf insteal of json / yaml / xml?
Protobuf is more lightweight and faster. This is because protobuf reprezents the data as an array of bytes. And, since the protobuf is strongly typed, the conversion will also be quite efficient since it won't use more than the number of bytes required for the data.

Also, unlike other data formats (like `json`, `yaml` or `xml`), `protobuf` is Language neutral: you can use the **schema**  defined in the proto file to compile it to any desired programming language. 