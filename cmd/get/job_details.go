package get

import (
	"encoding/json"
	"fmt"
	"github.com/adorigi/checkctl/pkg/output"
	"github.com/adorigi/checkctl/pkg/types"
	"io"
	"net/http"

	"github.com/adorigi/checkctl/pkg/config"
	"github.com/adorigi/checkctl/pkg/request"
	"github.com/adorigi/checkctl/pkg/utils"
	"github.com/spf13/cobra"
)

// jobDetailsCmd represents the controls command
var jobDetailsCmd = &cobra.Command{
	Use:   "job-details",
	Short: "Get Job Details",
	Long:  `Get Job Details`,
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

		jobId := utils.ReadStringFlag(cmd, "job-id")
		if jobId == "" {
			return fmt.Errorf("job-id flag is required")
		}

		var url string
		jobType := utils.ReadStringFlag(cmd, "job-type")
		switch jobType {
		case "compliance", "analytics", "discovery", "query":
			url = fmt.Sprintf("main/schedule/api/v3/job/%s/%s", jobType, jobId)
		default:
			return fmt.Errorf("please provide a valid job-type: compliance, analytics, discovery")
		}

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
		fmt.Println(string(body))

		if response.StatusCode != 200 {
			fmt.Println(string(body))
			return nil
		}

		switch jobType {
		case "query":
			var job types.GetAsyncQueryRunJobStatusResponse
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
		case "compliance":
			var job types.GetComplianceJobStatusResponse
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
		case "discovery":
			var job types.GetDescribeJobStatusResponse
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
		case "analytics":
			var job types.GetAnalyticsJobStatusResponse
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

		}

		return nil
	},
}
