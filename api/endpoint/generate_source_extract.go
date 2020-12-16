package endpoint

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/code"
	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/response"
	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/sourcemap"
)

type generateSourceExtractRequestBody struct {
	MapURL string `json:"map_url"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

type generateSourceExtractResponseBody struct {
	FileURL      string             `json:"file_url"`
	FunctionName string             `json:"function"`
	Line         int                `json:"line"`
	Column       int                `json:"column"`
	Extract      []code.LineExtract `json:"extract"`
}

func parseGenerateSourceExtractRequest(request events.APIGatewayProxyRequest) (
	*generateSourceExtractRequestBody,
	error,
) {
	body := generateSourceExtractRequestBody{}
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

// GenerateSourceExtract is the handler for POST /generate-source-extract
func GenerateSourceExtract(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	requestBody, err := parseGenerateSourceExtractRequest(request)
	if err != nil {
		return response.BadRequest(err)
	}

	fileURL, fn, line, column, extract, err := sourcemap.GenerateExtract(
		requestBody.MapURL,
		requestBody.Line,
		requestBody.Column,
	)

	if err != nil {
		return response.BadRequest(err)
	}
	return response.OK(generateSourceExtractResponseBody{
		FileURL:      fileURL,
		FunctionName: fn,
		Line:         line,
		Column:       column,
		Extract:      extract,
	})
}
