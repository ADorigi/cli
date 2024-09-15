package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List commands",
	Long:  `List commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}

func init() {

	ListCmd.PersistentFlags().Int("page-size", 25, "Defines page size of response")

	ListCmd.AddCommand(benchmarkSummaryCmd)
	benchmarkSummaryCmd.PersistentFlags().StringArray("integration-info", []string{}, "Integration info in the form 'integration=AWS,id=123,id_name=name'"+
		"values are optional and support regex")
	benchmarkSummaryCmd.PersistentFlags().String("benchmark-id", "", "Benchmark ID")
	benchmarkSummaryCmd.PersistentFlags().Int("top-integrations-count", 0, "Number of Top Integrations to show")

	ListCmd.AddCommand(jobsCmd)
	jobsCmd.PersistentFlags().StringArray("integration-info", []string{}, "Integration info in the form 'integration=AWS,id=123,id_name=name'"+
		"values are optional and support regex")
	jobsCmd.PersistentFlags().String("job-type", "", "Job Type: compliance/discovery/analytics")
	jobsCmd.PersistentFlags().String("selector", "", "Filter Type Selector: job-id/integration-info/status/benchmark")
	jobsCmd.PersistentFlags().StringSlice("job-id", []string{}, "Job ID")
	jobsCmd.PersistentFlags().StringSlice("status", []string{}, "Status")
	jobsCmd.PersistentFlags().StringSlice("benchmark", []string{}, "Benchmark")

}
