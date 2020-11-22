package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// # Your server-side functionality
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, World",
	}, nil

}
