# `first-friday-feb-2019`

Serverless Hack Day Project (Golang + OAuth)

## Instructions

Installing `serverless`

```
$ node --version
v10.15.0
$ npm --version
6.4.1
$ npm install
```

Creating the template and getting dependencies

```
$ go version
go version go1.11.4 darwin/amd64
$ ./node_modules/.bin/serverless create --template aws-go --path src/
$ go get github.com/aws/aws-lambda-go/events
```

Building the application and pushing it to AWS Lambda

```
$ cd src/
$ make
$ cd ..
$ AWS_ACCESS_KEY_ID=... \
>   AWS_SECRET_ACCESS_KEY=... \
>   serverless deploy
```

## Resources

- [Serverless][1] framework quickstart
- Serverless [Golang][2] instructions

## Oops (i.e. things that went wrong)

-   Visited `console.aws.amazon.com/lambda` for `us-east-2` (my account's
    default) rather than `us-east-1` (where our functions were deployed)
-   Got a 403 on the `hello` route (`@justinzhou93` modified the source while
    having build issues, this one may remain a mystery)

[1]: https://serverless.com/framework/docs/providers/aws/guide/quick-start/
[2]: https://serverless.com/blog/framework-example-golang-lambda-support/
