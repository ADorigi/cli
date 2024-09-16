/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/adorigi/checkctl/cmd/run"

	"github.com/adorigi/checkctl/cmd/get"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "checkctl",
	Short: "CLI for opengovernance",
	Long: `CLI for opengovernance
	Examples:

	Configure checkctl:
		checkctl configure 

	Run Discovery on one account:
		checkctl run discovery --integration id_name=account4
	
	Run discovery on all accounts:
    	checkctl run discovery --integration id_name=".*account.*"
	
	Run Discovery on all accounts with "prod" in name:
		checkctl run discovery --integration id_name=".*prod.*"

	Run Compliance Benchmark for a given benchmark ID:
		checkctl run compliance --benchmark-id aws_audit_manager_control_tower 

	Run Compliance Benchmark on a list of Integrations:
		checkctl run compliance --benchmark-id aws_audit_manager_control_tower --integration id_name="account1" --integration id_name="account2"
	
	Get Compliance Summary for an Integration
		go run . get compliance-summary-for-integration --benchmark-id aws_audit_manager_control_tower --integration id_name=account3
	
	Get Compliance Summary for a Benchmark
		go run . get compliance-summary-for-benchmark --benchmark-id aws_audit_manager_control_tower

	Get details of job with given id and type: 
		get job-details --job-id 301 --job-type compliance

	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(run.RunCmd)

	rootCmd.PersistentFlags().String("output", "", "Output format: json/table")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.checkctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
