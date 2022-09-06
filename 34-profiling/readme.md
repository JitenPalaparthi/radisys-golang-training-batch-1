```go test -cpuprofile cpu.prof -bench .```
```go tool pprof cpu.prof```
```top5 -cum```
- cumilative time taken

```go test -memprofile mem.prof -bench .```
	
```go test -cpuprofile cpu.prof -memprofile mem.prof -bench .```


