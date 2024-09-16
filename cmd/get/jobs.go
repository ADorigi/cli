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
	"strings"
)

// jobsCmd represents the List Jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Get Jobs for job type and in the time interval",
	Long:  `Get Jobs for job type and in the time interval`,
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
			fmt.Println("Error: must specify Job Type. Options: compliance, analytics, discovery")
			return nil
		}
		if strings.ToLower(jobType) != "compliance" && strings.ToLower(jobType) != "discovery" &&
			strings.ToLower(jobType) != "analytics" {
			fmt.Println("Invalid job type. Options: compliance, analytics, discovery")
			return nil
		}
		interval := utils.ReadStringFlag(cmd, "interval")
		if interval == "" {
			fmt.Println("Error: must specify interval like: 90m, 1h, 50 minutes, 2 hours")
			return nil
		}

		if err != nil {
			return err
		}
		url := fmt.Sprintf("main/schedule/api/v3/jobs/interval?job_type=%s&interval=%s", jobType, interval)
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

		var jobs []types.ListJobsByTypeItem
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
