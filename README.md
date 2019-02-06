# Getting Started

## Go

You need Go 1.6 or higher.

## gRPC

```sh
$ go get -u google.golang.org/grpc
```

## Protocol Buffers v3

- Download the binary (protoc-**version**-**os**) from their repo on [Github](https://github.com/protocolbuffers/protobuf/releases/latest).
- Unzip the file
- Add the files inside of the bin folder to your PATH

## Protobuf Go Plugin

```sh
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

- Add the plugin to your path too with:

```sh
$ export PATH=$PATH:$GOPATH/bin
```

# Clone this project

```sh
$ go get github.com/matAlmeida/hashlab-service-1-golang
$ cd $GOPATH/src/github.com/matAlmeida/hashlab-service-1-golang
$ cd product
$ protoc -I . product.proto --go_out=plugins=grpc:.
$ go run ../cmd/server
```

## [Link](https://github.com/hashlab/hiring/blob/master/challenges/pt-br/back-challenge.md)
