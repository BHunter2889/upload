build:
	gofmt -s -w ./
	go mod tidy -v
	go mod verify
	env GOOS=linux go build -ldflags="-s -w" -x

.PHONY: clean
clean:
	rm -f ~/bin/upload

.PHONY: deploy
deploy: clean build
	./upload local ~/bin upload
