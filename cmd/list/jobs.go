package list

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
	"strings"
)

// jobsCmd represents the List Jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
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

		jobType := utils.ReadStringFlag(cmd, "job-type")
		if jobType == "" {
			fmt.Println("Error: must specify Job Type")
			return nil
		}
		if strings.ToLower(jobType) != "compliance" && strings.ToLower(jobType) != "discovery" &&
			strings.ToLower(jobType) != "analytics" {
			fmt.Println("invalid job type")
			return nil
		}

		selector := utils.ReadStringFlag(cmd, "selector")
		if selector == "" {
			fmt.Println("Error: must specify Selector")
			return nil
		}
		fmt.Println(selector)
		if strings.ToLower(selector) != "job_id" && strings.ToLower(selector) != "integration_info" &&
			strings.ToLower(selector) != "status" && strings.ToLower(selector) != "benchmark" {
			fmt.Println("invalid selector. valid values: job_id, integration, status, benchmark")
			return nil
		}

		integrationsStr, err := utils.ReadStringArrayFlag(cmd, "integration")
		if err != nil {
			return err
		}

		var integrations []types.IntegrationFilterInfo
		for _, integrationStr := range integrationsStr {
			integrations = append(integrations, types.ParseIntegrationInfo(integrationStr))
		}

		jobId, err := utils.ReadStringSliceFlag(cmd, "job-id")
		if err != nil {
			return err
		}
		status, err := utils.ReadStringSliceFlag(cmd, "status")
		if err != nil {
			return err
		}
		benchmark, err := utils.ReadStringSliceFlag(cmd, "benchmark")
		if err != nil {
			return err
		}

		requestPayload := types.ListJobsByTypeRequest{
			JobType:         jobType,
			Selector:        selector,
			JobId:           jobId,
			IntegrationInfo: integrations,
			Status:          status,
			Benchmark:       benchmark,
		}

		payload, err := json.Marshal(requestPayload)
		if err != nil {
			return err
		}
		url := fmt.Sprintf("main/schedule/api/v2/jobs")
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

		var jobs []types.ListJobsByTypeResponse
		err = json.Unmarshal(body, &jobs)
		if err != nil {
			return err
		}

		if outputFormat == "table" {
			fmt.Println("Table view not supported, use json view: --output json")
			// TODO
		} else {
			js, err := json.MarshalIndent(jobs, "", "   ")
			if err != nil {
				return err
			}
			fmt.Print(string(js))
		}

		return nil
	},
}
