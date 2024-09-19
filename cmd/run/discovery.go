package run

import (
	"encoding/json"
	"fmt"
	"github.com/adorigi/checkctl/pkg/output"
	"io"
	"net/http"

	"github.com/adorigi/checkctl/pkg/output/tables"
	"github.com/adorigi/checkctl/pkg/request"

	"github.com/adorigi/checkctl/pkg/config"
	"github.com/adorigi/checkctl/pkg/types"
	"github.com/adorigi/checkctl/pkg/utils"
	"github.com/spf13/cobra"
)

// benchmarksCmd represents the benchmarks command
var discoveryCmd = &cobra.Command{
	Use:   "discovery",
	Short: "Run specified benchmark on given integrations",
	Long:  `Run specified benchmark on given integrations.`,
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

		integrationsStr, err := utils.ReadStringArrayFlag(cmd, "integration")
		if err != nil {
			return err
		}
		resourceTypes, err := utils.ReadStringSliceFlag(cmd, "resource-type")
		if err != nil {
			return err
		}
		forceFull := utils.ReadBoolFlag(cmd, "force-full")

		var integrations []types.IntegrationFilterInfo
		for _, integrationStr := range integrationsStr {
			if _, ok := configuration.Integrations[integrationStr]; ok {
				fmt.Printf("Found stored integration %s", integrationStr)
				integrationStr = configuration.Integrations[integrationStr]
			}
			integrations = append(integrations, types.ParseIntegrationInfo(integrationStr))
		}
		req := types.RunDiscoveryRequest{
			IntegrationInfo: integrations,
			ResourceTypes:   resourceTypes,
			ForceFull:       forceFull,
		}

		payload, err := json.Marshal(req)
		if err != nil {
			return err
		}

		url := fmt.Sprintf("main/schedule/api/v3/discovery/run")
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

		var runDiscoveryResponse []types.RunDiscoveryResponse
		err = json.Unmarshal(body, &runDiscoveryResponse)
		if err != nil {
			return err
		}

		if outputFormat == "table" {
			rows, err := utils.GenerateDiscoveryJobsRows(runDiscoveryResponse)
			if err != nil {
				return err
			}

			tables.PrintDiscoveryJobsTable(rows)
		} else {
			return output.OutputJson(cmd, runDiscoveryResponse)
		}

		return nil
	},
}
