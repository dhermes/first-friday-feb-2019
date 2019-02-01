.PHONY: help build clean deploy update-gimme-key update-hello update-who-am-i remove

help:
	@echo 'Makefile for `first-friday-feb-2019` app'
	@echo ''
	@echo 'Usage:'
	@echo '   make build               Build the `go` application'
	@echo '   make clean               Remove the built application'
	@echo '   make deploy              Deploy the built application to AWS Lambda'
	@echo '   make update-gimme-key    Update the (already deployed) `gimme-key` function'
	@echo '   make update-hello        Update the (already deployed) `hello` function'
	@echo '   make update-who-am-i     Update the (already deployed) `who-am-i` function'
	@echo '   make remove              Remove the AWS Lambda service'
	@echo ''
	@echo 'Set the AWS_SECRET_ACCESS_KEY and AWS_ACCESS_KEY_ID variables for deploy'
	@echo ''

build:
	[ -d ./node_modules ] || npm ci
	env GOOS=linux go build -ldflags="-s -w" -o src/bin/gimme-key src/gimme-key/main.go
	env GOOS=linux go build -ldflags="-s -w" -o src/bin/hello src/hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o src/bin/who-am-i src/who-am-i/main.go

clean:
	rm -rf ./src/bin

deploy: clean build
	cd src/ && ../node_modules/.bin/serverless deploy --verbose

update-gimme-key: clean build
	cd src/ && ../node_modules/.bin/serverless deploy function --function gimme-key

update-hello: clean build
	cd src/ && ../node_modules/.bin/serverless deploy function --function hello

update-who-am-i: clean build
	cd src/ && ../node_modules/.bin/serverless deploy function --function who-am-i

remove:
	cd src/ && ../node_modules/.bin/serverless remove
