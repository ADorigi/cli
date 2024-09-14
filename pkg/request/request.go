package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adorigi/opengovernance/pkg/types"
	"github.com/adorigi/opengovernance/pkg/utils"
	"github.com/spf13/cobra"
)

func GenerateRequest(
	apiKey string,
	apiEndpoint string,
	cmd *cobra.Command,
	requestMethod string,
	requestPath string,
) (*http.Request, error) {

	bearer := fmt.Sprintf("Bearer %s", apiKey)

	requestPayload := types.RequestPayload{
		Cursor:  1,
		PerPage: int(utils.ReadIntFlag(cmd, "page-size")),
	}

	payload, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s/%s", apiEndpoint, requestPath)
	request, err := http.NewRequest(requestMethod, uri, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", bearer)
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
