package response

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// JSON returns a json-encoded response
func JSON(data interface{}, statusCode int) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer
	body, err := json.Marshal(&data)
	if err != nil {
		panic(err)
	}
	json.HTMLEscape(&buf, body)
	resp := events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}
