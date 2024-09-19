package output

import (
	"encoding/json"
	"github.com/spf13/cobra"
)

func OutputJson(cmd *cobra.Command, result any) error {
	encoder := json.NewEncoder(cmd.OutOrStdout())
	encoder.SetIndent("", "   ")
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(result)
	if err != nil {
		return err
	}
	return nil
}
