- To run go
```go run hello.go```
- To know where binary/exe created when run
```go run --work hello.go```
- To build and gen a binary/exe
```go build hello.go```
- To build and generate desired output binary/exe
```go build -o app hello.go```
- To build for other operating systems
```GOOS=linux GOARCH=amd64 go build -o app hello.go```
- To know list of OS and ARCH
```go tool dist list```
- To install binaries/exe to GOBIN location
```go install hello.go```
- To install binaries/exe from remote repositories
```go install github.com/JitenPalaparthi/urllinter@latest```
- To compile and link
```go tool compile hello.go```
```go tool like -o hello.exe hello.o```
- To go mod
- 1- standard packages       ---> GOROOT -> src/
- 2- user defined packages   ---> GOPATH -> pkg/ To setup GOPATH -> src/ pkg/ bin/ < go 1.11
- 3- third party packages    ---> 
- go mod --> 1. To work with user defined packages
-            2. To maintain dependencies
```go mod init demo```