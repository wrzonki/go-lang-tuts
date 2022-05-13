# Usage
```bash
go mod init github.com/wrzonki/go-lang-tuts
go run main.go
go build main.go
```

# Unit tests
```bash
go test
go test ./package-name
go test -v
go test -v -run=TestAdd
go test -v -run='TestAdd|TestMultiply'
go test -v -run='TestMultiply/2*3=6'
go test -cover
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```