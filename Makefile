build:
	gofmt -s -w ./
	go mod tidy -v
	go mod verify
	env GOOS=linux go build -ldflags="-s -w" -x