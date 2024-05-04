package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/oosawy/imageon"
)

func handleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	res, err := imageon.HandleRequest(ctx, imageon.RawRequest{
		Path: request.RawPath,
		Body: request.Body,
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: res.StatusCode,
		Body:       res.Body,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
