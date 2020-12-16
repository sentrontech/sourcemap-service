package endpoint

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/response"
	"github.com/jpstevens/sentron-sourcemaps/internal/pkg/sourcemap"
)

type locateSourcemapRequestBody struct {
	URL string `json:"url"`
}

type locateSourcemapResponseBody struct {
	MapURL  string `json:"map_url"`
	IsGuess bool   `json:"is_guess"`
}

func parseLocateSourcemapRequest(request events.APIGatewayProxyRequest) (
	*locateSourcemapRequestBody,
	error,
) {
	body := locateSourcemapRequestBody{}
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

// LocateSourcemap is the handler for POST /locate-sourcemap
func LocateSourcemap(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	requestBody, err := parseLocateSourcemapRequest(request)
	if err != nil {
		return response.BadRequest(err)
	}

	sourceMapURL, isGuess, err := sourcemap.Locate(requestBody.URL)

	if err != nil {
		return response.BadRequest(err)
	}
	return response.OK(locateSourcemapResponseBody{
		MapURL:  sourceMapURL,
		IsGuess: isGuess,
	})
}
