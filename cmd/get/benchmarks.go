/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adorigi/checkctl/pkg/output/tables"

	"github.com/adorigi/checkctl/pkg/config"
	"github.com/adorigi/checkctl/pkg/request"
	"github.com/adorigi/checkctl/pkg/types"
	"github.com/adorigi/checkctl/pkg/utils"
	"github.com/spf13/cobra"
)

// benchmarksCmd represents the benchmarks command
var benchmarksCmd = &cobra.Command{
	Use:   "benchmarks",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		requestPayload := types.GetBenchmarkPayload{
			Cursor:                 int(utils.ReadIntFlag(cmd, "page-number")),
			PerPage:                int(utils.ReadIntFlag(cmd, "page-size")),
			OnlyRootBenchmark:      utils.ReadBoolFlag(cmd, "show-only-root"),
			IncludeFindingsSummary: utils.ReadBoolFlag(cmd, "include-findings-summary"),
		}

		payload, err := json.Marshal(requestPayload)
		if err != nil {
			return err
		}

		request, err := request.GenerateRequest(
			configuration.ApiKey,
			configuration.ApiEndpoint,
			"POST",
			"main/compliance/api/v3/benchmarks",
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

		var getBenchmarksResponse types.GetBenchmarksResponse
		err = json.Unmarshal(body, &getBenchmarksResponse)
		if err != nil {
			return err
		}

		if outputFormat == "table" {
			rows := utils.GenerateBenchmarkRows(getBenchmarksResponse.Items)

			tables.PrintBenchmarksTable(rows)
		} else {
			js, err := json.MarshalIndent(getBenchmarksResponse.Items, "", "   ")
			if err != nil {
				return err
			}
			fmt.Print(string(js))
		}

		fmt.Printf(
			"\n\n\n\nNext Page: \n\tcheckctl get benchmarks --page-size %d --page-number %d --output %s\n",
			utils.ReadIntFlag(cmd, "page-size"),
			utils.ReadIntFlag(cmd, "page-number")+1,
			outputFormat,
		)

		return nil
	},
}

func init() {
	benchmarksCmd.Flags().Bool("show-only-root", true, "Show only root benchmarks(default: true)")
	benchmarksCmd.Flags().Bool("include-findings-summary", false, "Include findings summary in response(default: false)")

}
