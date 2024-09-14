clean:
	rm -r ./bin/

build: clean
	go build -o bin/
	env GOOS=linux GOARCH=arm go build -o bin/opengovernance-linux

goreleaser:
	REPOSITORY_OWNER=local REPOSITORY_NAME=local goreleaser build --snapshot