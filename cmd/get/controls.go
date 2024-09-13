/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adorigi/opengovernance/pkg/config"
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
		bearer := fmt.Sprintf("Bearer %s", configuration.ApiKey)
		// payload := []byte(`
		// {
		// 	"cursor": 1,
		// 	"per_page": 10
		// }
		// `)

		requestPayload := types.RequestPayload{
			Cursor:  1,
			PerPage: int(utils.ReadIntFlag(cmd, "page-size")),
		}

		payload, err := json.Marshal(requestPayload)
		if err != nil {
			return err
		}

		request, err := http.NewRequest("GET", "https://demo4.kaytu.sh/main/compliance/api/v2/controls", bytes.NewBuffer(payload))
		if err != nil {
			return err
		}
		request.Header.Add("Authorization", bearer)
		request.Header.Set("Content-Type", "application/json")

		response, err := client.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		var controls []types.Control
		err = json.Unmarshal(body, &controls)
		if err != nil {
			return err
		}
		fmt.Println(len(controls))

		return nil
	},
}

func init() {

}
