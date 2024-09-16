/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/adorigi/checkctl/pkg/config"
	"github.com/adorigi/checkctl/pkg/input"
	"github.com/adorigi/checkctl/pkg/utils"
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

		var configuration *config.Configuration
		var err error

		if cmd.Flags().Changed("output") && cmd.Flags().Changed("app-endpoint") && cmd.Flags().Changed("api-key") {

			configuration = config.NewConfiguration(
				utils.ReadStringFlag(cmd, "output"),
				utils.ReadStringFlag(cmd, "app-endpoint"),
				utils.ReadStringFlag(cmd, "utilization-analyzer-endpoint"),
				utils.ReadStringFlag(cmd, "api-key"),
			)

		} else {
			configuration, err = input.GetConfigurationFromForm()
			if err != nil {
				return err
			}
		}

		err = config.CreateConfigFile(configuration)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {

	configureCmd.Flags().String("output", "", "Output format")
	configureCmd.Flags().String("app-endpoint", "", "App endpoint for API")
	configureCmd.Flags().String("utilization-analyzer-endpoint", "https://optimizer.kaytu.io/", "Endpoint for Utilization and Optimization Service")
	configureCmd.Flags().String("api-key", "", "API key")

}
