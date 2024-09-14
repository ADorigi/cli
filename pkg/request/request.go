package request

import (
	"bytes"
	"fmt"
	"net/http"
)

func GenerateRequest(
	apiKey string,
	apiEndpoint string,
	requestMethod string,
	requestPath string,
	payload []byte,
) (*http.Request, error) {

	bearer := fmt.Sprintf("Bearer %s", apiKey)

	uri := fmt.Sprintf("%s/%s", apiEndpoint, requestPath)
	request, err := http.NewRequest(requestMethod, uri, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
