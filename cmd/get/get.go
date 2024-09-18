/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: checkctl get controls|benchmarks --page-size")
	},
}

func init() {

	GetCmd.PersistentFlags().Int("page-size", 25, "Defines page size of response")
	GetCmd.PersistentFlags().Int("page-number", 1, "Defines page number of response")

	controlsCmd.Flags().StringSlice("benchmark-id", []string{}, "List of Benchmark IDs to get the description for")

	GetCmd.AddCommand(controlsCmd)
	GetCmd.AddCommand(benchmarksCmd)

	GetCmd.AddCommand(complianceSummaryForIntegrationCmd)
	complianceSummaryForIntegrationCmd.PersistentFlags().String("integration", "", "Integration info in the form 'integration=AWS,id=123,id_name=name'"+
		"values are optional and support regex")
	complianceSummaryForIntegrationCmd.PersistentFlags().String("benchmark-id", "", "Benchmark ID")

	GetCmd.AddCommand(complianceSummaryForBenchmarkCmd)
	complianceSummaryForBenchmarkCmd.PersistentFlags().StringSlice("benchmark-id", []string{}, "List of Benchmark IDs to get the summary for (optional)")
	complianceSummaryForBenchmarkCmd.PersistentFlags().Bool("is-root", true, "Whether to return only root benchmarks or not. (matters if benchmark-id list not provided)")

	GetCmd.AddCommand(jobDetailsCmd)
	jobDetailsCmd.PersistentFlags().String("job-id", "", "Job ID")
	jobDetailsCmd.PersistentFlags().String("job-type", "", "Job Type. Options: compliance, analytics, discovery")

	GetCmd.AddCommand(jobsCmd)
	jobsCmd.PersistentFlags().String("job-type", "", "Job Type. Options: compliance, analytics, discovery")
	jobsCmd.PersistentFlags().String("interval", "90m", "Specify time interval like: 90m, 1h, 50 minutes, 2 hours")

	GetCmd.AddCommand(findingsCmd)
	findingsCmd.PersistentFlags().StringArray("integration", []string{}, "Integration info in the form 'integration=AWS,id=123,id_name=name'"+
		"values are optional and support regex")
	findingsCmd.PersistentFlags().StringSlice("benchmark-id", []string{}, "Benchmark ID")
}
