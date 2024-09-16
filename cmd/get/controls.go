/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adorigi/checkctl/pkg/config"
	"github.com/adorigi/checkctl/pkg/request"
	"github.com/adorigi/checkctl/pkg/types"
	"github.com/adorigi/checkctl/pkg/utils"
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

		requestPayload := types.RequestPayload{
			Cursor:  int(utils.ReadIntFlag(cmd, "page-number")),
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
			"main/compliance/api/v3/controls",
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

		var getControlsResponse types.GetControlsResponse
		err = json.Unmarshal(body, &getControlsResponse)
		if err != nil {
			return err
		}

		// if outputFormat == "table" {
		// 	rows := utils.GenerateControlRows(getControlsResponse.Items)

		// 	tables.PrintControlsTable(rows)
		// } else {
		js, err := json.MarshalIndent(getControlsResponse.Items, "", "   ")
		if err != nil {
			return err
		}
		fmt.Print(string(js))
		// }

		fmt.Printf(
			"\n\n\n\nNext Page: \n\tcheckctl get controls --page-size %d --page-number %d --output %s\n",
			utils.ReadIntFlag(cmd, "page-size"),
			utils.ReadIntFlag(cmd, "page-number")+1,
			outputFormat,
		)

		return nil
	},
}
