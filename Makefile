.PHONY: help build clean deploy update-hello clean-up

help:
	@echo 'Makefile for `first-friday-feb-2019` app'
	@echo ''
	@echo 'Usage:'
	@echo '   make build           Build the `go` application'
	@echo '   make clean           Remove the built application'
	@echo '   make deploy          Deploy the built application to AWS Lambda'
	@echo '   make update-hello    Update the (already deployed) `hello` function'
	@echo '   make clean-up        Remove the AWS Lambda service'
	@echo ''
	@echo 'Set the AWS_SECRET_ACCESS_KEY and AWS_ACCESS_KEY_ID variables for deploy'
	@echo ''

build:
	[ -d ./node_modules ] || npm ci
	env GOOS=linux go build -ldflags="-s -w" -o src/bin/hello src/hello/main.go

clean:
	rm -rf ./src/bin

deploy: clean build
	cd src/ && ../node_modules/.bin/serverless deploy --verbose

update-hello: clean build
	cd src/ && ../node_modules/.bin/serverless deploy function --function hello

clean-up:
	cd src/ && ../node_modules/.bin/serverless remove
