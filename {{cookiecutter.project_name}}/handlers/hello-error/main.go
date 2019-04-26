package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context) (string, error) {
	return "", errors.New("goodbye")
}
