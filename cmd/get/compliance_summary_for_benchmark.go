package get

import (
	"encoding/json"
	"fmt"
	"github.com/adorigi/checkctl/pkg/config"
	"github.com/adorigi/checkctl/pkg/request"
	"github.com/adorigi/checkctl/pkg/types"
	"github.com/adorigi/checkctl/pkg/utils"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

// complianceSummaryForBenchmarkCmd represents the benchmarks command
var complianceSummaryForBenchmarkCmd = &cobra.Command{
	Use:   "compliance-summary-for-benchmark",
	Short: "Get compliance summary for benchmark",
	Long:  `Get compliance summary for benchmark`,
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

		benchmarkIDs, err := utils.ReadStringSliceFlag(cmd, "benchmark-id")
		if err != nil {
			return err
		}

		isRoot := utils.ReadBoolFlag(cmd, "is-root")
		requestPayload := types.ComplianceSummaryOfBenchmarkRequest{
			Benchmarks: benchmarkIDs,
			IsRoot:     &isRoot,
		}

		payload, err := json.Marshal(requestPayload)
		if err != nil {
			return err
		}

		url := fmt.Sprintf("main/compliance/api/v3/compliance/summary/benchmark")
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

		var summary []types.ComplianceSummaryOfBenchmarkResponse
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
