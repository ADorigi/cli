package list

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/adorigi/opengovernance/pkg/config"
	"github.com/adorigi/opengovernance/pkg/request"
	"github.com/adorigi/opengovernance/pkg/types"
	"github.com/adorigi/opengovernance/pkg/utils"
	"github.com/spf13/cobra"
)

// benchmarksCmd represents the benchmarks command
var benchmarkSummaryCmd = &cobra.Command{
	Use:   "benchmark-summary",
	Short: "Get benchmark findings summary",
	Long:  `Get benchmark findings summary`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{}
		configuration, err := config.ReadConfigFile()
		if err != nil {
			return err
		}

		outputFormat := utils.ReadStringFlag(cmd, "output")
		if outputFormat == "" {
			outputFormat = configuration.OutputFormat
		}

		benchmarkId := utils.ReadStringFlag(cmd, "benchmark-id")
		if benchmarkId == "" {
			fmt.Println("Error: must specify benchmark id")
			return nil
		}

		integrationsStr, err := utils.ReadStringArrayFlag(cmd, "integration-info")
		if err != nil {
			return err
		}

		var integrations []types.IntegrationFilterInfo
		for _, integrationStr := range integrationsStr {
			integrations = append(integrations, types.ParseIntegrationInfo(integrationStr))
		}

		var topIntegrationsCount int
		topIntegrationsCountStr := utils.ReadStringFlag(cmd, "top-integrations-count")
		if topIntegrationsCountStr == "" {
			topIntegrationsCount = 0
		} else {
			topIntegrationsCount, err = strconv.Atoi(topIntegrationsCountStr)
		}

		requestPayload := types.GetBenchmarkSummaryV2Request{
			Integration:          integrations,
			TopIntegrationsCount: topIntegrationsCount,
		}

		payload, err := json.Marshal(requestPayload)
		if err != nil {
			return err
		}
		url := fmt.Sprintf("main/compliance/api/v2/benchmark/%s/summary", benchmarkId)
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

		var summary types.GetBenchmarkSummaryV2Response
		err = json.Unmarshal(body, &summary)
		if err != nil {
			return err
		}

		if outputFormat == "table" {
			fmt.Println("Table view not supported, use json view: --output json")
			// TODO
		} else {
			js, err := json.MarshalIndent(summary, "", "   ")
			if err != nil {
				return err
			}
			fmt.Print(string(js))
		}

		return nil
	},
}
