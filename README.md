# aerospike-set-deleter
A handy tool to delete all records in a set of Aerospike

## Install
Just go get a binary from [releases](https://github.com/okuramasafumi/aerospike-set-deleter/releases) and use it.
Alternatively, you can run `go get github.com/okuramasafumi/aerospike-set-deleter` and install it yourself.

## Usage
```sh
aerospike-set-deleter -host 127.0.0.1 -port 3000 -namespace ns -set foo
```

`-set` parameter is required. Other parameters are optional (you can see the defaults by just running the binary without any options).
