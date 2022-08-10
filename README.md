- To run go
```go run hello.go```
- To build and gen a binary/exe
```go build hello.go```
- To build and generate desired output binary/exe
```go build -o app hello.go```
- To build for other operating systems
```GOOS=linux GOARCH=amd64 go build -o app hello.go```
- To know list of OS and ARCH
```go tool dist list```