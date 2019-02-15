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
$ go get github.com/aws/aws-sdk-go/aws
$ go get github.com/dgrijalva/jwt-go
$ go get github.com/satori/go.uuid
$ # Wrong way to do this, but "meh" (Justin will figure it out).
$ go get github.com/dhermes/first-friday-feb-2019/pkg/verify
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

## Application Design

We have 3 pieces to build

-   Lambda Function `/gimme-key` which takes an (unauthorized) request and
    returns a service account name and a private key (then stores the
    corresponding public key in S3 or DynamoDB)
-   Lambda Function `/who-am-i` which requires a JWT for authentication (maybe
    using a `scope` claim for authorization) and just returns some metadata
    about the authorized user
-   Client which calls `/gimme-key` to get a key and then creates an
    authenaticated request to `/who-am-i` with that key

## Documentation

-   [Serverless][1] framework quickstart
-   Serverless [Golang][2] instructions

## Oops (i.e. things that went wrong)

-   Visited `console.aws.amazon.com/lambda` for `us-east-2` (my account's
    default) rather than `us-east-1` (where our functions were deployed)
-   Got a 403 on the `hello` route (`@justinzhou93` modified the source while
    having build issues, this one may remain a mystery)

[1]: https://serverless.com/framework/docs/providers/aws/guide/quick-start/
[2]: https://serverless.com/blog/framework-example-golang-lambda-support/
