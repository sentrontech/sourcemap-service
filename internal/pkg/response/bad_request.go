package response

import (
	"github.com/aws/aws-lambda-go/events"
)

type badRequestBody struct {
	Message string `json:"message"`
}

// BadRequest returns a 400 error JSON response
func BadRequest(err error) (events.APIGatewayProxyResponse, error) {
	return JSON(badRequestBody{
		Message: err.Error(),
	}, 400)
}
