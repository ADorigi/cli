package get

import (
	"encoding/json"
	"fmt"
	"github.com/adorigi/checkctl/pkg/output"
	"github.com/adorigi/checkctl/pkg/output/tables"
	"github.com/adorigi/checkctl/pkg/types"
	"io"
	"net/http"

	"github.com/adorigi/checkctl/pkg/config"
	"github.com/adorigi/checkctl/pkg/request"
	"github.com/adorigi/checkctl/pkg/utils"
	"github.com/spf13/cobra"
)

// queryResultCmd represents the controls command
var queryResultCmd = &cobra.Command{
	Use:   "query-result",
	Short: "Get async query run result by run id",
	Long:  `Get async query run result by run id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{}
		configuration, err := config.ReadConfigFile()
		if err != nil {
			return err
		}

		outputFormat := utils.ReadStringFlag(cmd, "output")
		if outputFormat == "" {
			outputFormat = "table"
		}

		runId := utils.ReadStringFlag(cmd, "run-id")
		if runId == "" {
			return fmt.Errorf("run-id flag is required")
		}

		var url string

		url = fmt.Sprintf("main/inventory/api/v3/query/async/run/%s/result", runId)

		request, err := request.GenerateRequest(
			configuration.ApiKey,
			configuration.ApiEndpoint,
			"GET",
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
		var result types.GetAsyncQueryRunResultResponse

		err = json.Unmarshal(body, &result)
		if err != nil {
			return err
		}

		if outputFormat == "table" {
			tables.PrintQueryResultTable(result.ColumnNames, result.Result)
		} else {
			return output.OutputJson(cmd, result)
		}

		return nil
	},
}
