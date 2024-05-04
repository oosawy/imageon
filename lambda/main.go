package main

import (
	"github.com/oosawy/imageon"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(imageon.HandleRequest)
}
