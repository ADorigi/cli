clean:
	rm -r ./bin/

build: clean
	go build -o bin/
	env GOOS=linux GOARCH=arm go build -o bin/checkctl-linux

goreleaser:
	REPOSITORY_OWNER=local REPOSITORY_NAME=local goreleaser build --snapshot

install: 
	rm $(GOPATH)/bin/checkctl
	go install