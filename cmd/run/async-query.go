package run

import (
	"encoding/json"
	"fmt"
	"github.com/adorigi/checkctl/pkg/output"
	"io"
	"net/http"

	"github.com/adorigi/checkctl/pkg/request"

	"github.com/adorigi/checkctl/pkg/config"
	"github.com/adorigi/checkctl/pkg/types"
	"github.com/adorigi/checkctl/pkg/utils"
	"github.com/spf13/cobra"
)

// benchmarksCmd represents the benchmarks command
var asyncQueryCmd = &cobra.Command{
	Use:   "async-query",
	Short: "Run specified query (control or named-query) async",
	Long:  `Run specified query (control or named-query) async by given query id`,
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

		queryId := utils.ReadStringFlag(cmd, "query-id")
		if err != nil {
			return err
		}

		if queryId == "" {
			return fmt.Errorf("query id required")
		}

		url := fmt.Sprintf("main/schedule/api/v3/query/%s/run", queryId)
		request, err := request.GenerateRequest(
			configuration.ApiKey,
			configuration.ApiEndpoint,
			"PUT",
			url,
			[]byte{},
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

		var job types.QueryRunnerJob
		err = json.Unmarshal(body, &job)
		if err != nil {
			return err
		}

		if outputFormat == "table" {
			fmt.Println("Table view not supported, use json view: --output json")
			//TODO
		} else {
			return output.OutputJson(cmd, job)
		}

		return nil
	},
}
