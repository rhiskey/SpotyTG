```go
go test -cover
```
### Generating an HTML coverage report
```go
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html 
```