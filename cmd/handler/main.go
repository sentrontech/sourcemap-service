package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jpstevens/sentron-sourcemaps/api/endpoint"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.Path {
	case "/locate-sourcemap":
		return endpoint.LocateSourcemap(request)
	case "/generate-source-extract":
		return endpoint.GenerateSourceExtract(request)
	default:
		panic(fmt.Errorf("Unknown path"))
	}
}

func main() {
	lambda.Start(handler)
}
