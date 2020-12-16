package response

import (
	"github.com/aws/aws-lambda-go/events"
)

type serverErrorBody struct {
	Message string `json:"message"`
}

// ServerError returns a 500 error JSON response
func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	return JSON(serverErrorBody{
		Message: err.Error(),
	}, 500)
}
