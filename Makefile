build:
	echo "Compiling for Linux:"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o bin/mygopass main.go

buildall:
	echo "Compiling for every OS and Platform"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o bin/mygopass-linux main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=freebsd go build -a -o bin/mygopass-freebsd main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -a -o bin/mygopass-windows main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -a -o bin/mygopass-darwin main.go

run:
	go run main.go
