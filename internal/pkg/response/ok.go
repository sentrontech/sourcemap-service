package response

import (
	"github.com/aws/aws-lambda-go/events"
)

// OK returns a 200 JSON response
func OK(data interface{}) (events.APIGatewayProxyResponse, error) {
	return JSON(data, 200)
}
