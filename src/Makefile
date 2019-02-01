.PHONY: build clean deploy update-hello update-world clean-up

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go

clean:
	rm -rf ./bin

deploy: clean build
    serverless deploy --verbose

update-hello: clean build
    serverless deploy function --function hello

update-world: clean build
    serverless deploy function --function world

clean-up:
    serverless remove
