/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/adorigi/cli/pkg/config"
	"github.com/adorigi/cli/pkg/utils"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configuraton := config.NewConfiguration(
			utils.ReadStringFlag(cmd, "output"),
			utils.ReadStringFlag(cmd, "endpoint"),
			utils.ReadStringFlag(cmd, "apikey"),
		)

		err := config.CreateConfigFile(configuraton)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {

	configureCmd.Flags().String("output", "", "Output format")
	configureCmd.Flags().String("endpoint", "", "Endpoint for API")
	configureCmd.Flags().String("apikey", "", "API key")

	configureCmd.MarkFlagRequired("output")
	configureCmd.MarkFlagRequired("endpoint")
	configureCmd.MarkFlagRequired("apikey")

}
