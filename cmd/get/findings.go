package get

import (
	"encoding/json"
	"fmt"
	"github.com/adorigi/checkctl/pkg/output"
	"io"
	"net/http"

	"github.com/adorigi/checkctl/pkg/config"
	"github.com/adorigi/checkctl/pkg/request"
	"github.com/adorigi/checkctl/pkg/types"
	"github.com/adorigi/checkctl/pkg/utils"
	"github.com/spf13/cobra"
)

var findingsCmd = &cobra.Command{
	Use:   "findings",
	Short: "Get findings with the given filters",
	Long:  `Get findings with the given filters`,
	RunE: func(cmd *cobra.Command, args []string) error {

		client := &http.Client{}
		configuration, err := config.ReadConfigFile()
		if err != nil {
			return err
		}

		// outputFormat := utils.ReadStringFlag(cmd, "output")
		// if outputFormat == "" {
		// 	outputFormat = configuration.OutputFormat
		// }

		benchmarkIDs, err := utils.ReadStringSliceFlag(cmd, "benchmark-id")
		if err != nil {
			return err
		}

		if _, ok := configuration.Benchmarks[benchmarkIDs[0]]; ok {
			fmt.Printf("Found stored Benchmark IDs %s", benchmarkIDs[0])
			benchmarkIDs = configuration.Benchmarks[benchmarkIDs[0]]
		}

		integrationStr := utils.ReadStringFlag(cmd, "integration")
		if integrationStr == "" {
			fmt.Println(`Error: must specify integration
				Integration info in the form 'integration=AWS,id=123,id_name=name'`)
			return nil
		}

		integrationsStr, err := utils.ReadStringArrayFlag(cmd, "integration")
		if err != nil {
			return err
		}

		var integrations []types.IntegrationFilterInfo
		for _, integrationStr := range integrationsStr {
			if _, ok := configuration.Integrations[integrationStr]; ok {
				fmt.Printf("Found stored integration %s", integrationStr)
				integrationStr = configuration.Integrations[integrationStr]
			}
			integrations = append(integrations, types.ParseIntegrationInfo(integrationStr))
		}

		requestPayload := types.FindingsRequestPayload{
			Filters: types.FindingRequestFilters{
				BenchmarkIDs: benchmarkIDs,
				Integrations: integrations,
			},
		}

		payload, err := json.Marshal(requestPayload)
		if err != nil {
			return err
		}

		url := "main/compliance/api/v3/findings"
		request, err := request.GenerateRequest(
			configuration.ApiKey,
			configuration.ApiEndpoint,
			"POST",
			url,
			payload,
		)
		if err != nil {
			return err
		}

		response, err := client.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		if response.StatusCode != 200 {
			fmt.Println(string(body))
			return nil
		}

		var findingsResponse types.FindingsResponse
		err = json.Unmarshal(body, &findingsResponse)
		if err != nil {
			return err
		}

		return output.OutputJson(cmd, findingsResponse)
	},
}
