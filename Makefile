run:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/cmd . && \
	./bin/cmd