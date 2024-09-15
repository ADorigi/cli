/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adorigi/opengovernance/pkg/output/tables"

	"github.com/adorigi/opengovernance/pkg/config"
	"github.com/adorigi/opengovernance/pkg/request"
	"github.com/adorigi/opengovernance/pkg/types"
	"github.com/adorigi/opengovernance/pkg/utils"
	"github.com/spf13/cobra"
)

// controlsCmd represents the controls command
var controlsCmd = &cobra.Command{
	Use:   "controls",
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

		var cursor = 1
		var controls []types.Control

		for {

			fmt.Println(cursor)

			requestPayload := types.RequestPayload{
				Cursor:  cursor,
				PerPage: int(utils.ReadIntFlag(cmd, "page-size")),
			}

			payload, err := json.Marshal(requestPayload)
			if err != nil {
				return err
			}

			request, err := request.GenerateRequest(
				configuration.ApiKey,
				configuration.ApiEndpoint,
				"POST",
				"main/compliance/api/v2/controls",
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

			var getControlsResponse types.GetControlsResponse
			err = json.Unmarshal(body, &getControlsResponse)
			if err != nil {
				return err
			}

			if len(getControlsResponse.Items) == 0 {
				break
			}

			controls = append(controls, getControlsResponse.Items...)
			cursor += 1

		}

		if outputFormat == "table" {
			rows := utils.GenerateControlRows(controls)

			tables.PrintControlsTable(rows)
		} else {
			js, err := json.MarshalIndent(controls, "", "   ")
			if err != nil {
				return err
			}
			fmt.Print(string(js))
		}

		return nil
	},
}
