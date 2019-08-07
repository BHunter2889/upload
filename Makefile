build:
	gofmt -s -w ./
	go get
	env GOOS=linux go build -ldflags="-s -w" -x -o upload