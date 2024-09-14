package run

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RunCmd represents the get command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run discovery or benchmark jobs",
	Long:  `Run discovery or benchmark jobs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: opengovernance run discovery|benchmark")
	},
}

func init() {

	RunCmd.AddCommand(benchmarkCmd)

	benchmarkCmd.PersistentFlags().StringArray("integration-info", []string{}, "Integration info in the form 'integration=AWS,id=123,id_name=name'"+
		"values are optional and support regex")
	benchmarkCmd.PersistentFlags().String("benchmark-id", "", "Benchmark ID")

	RunCmd.AddCommand(discoveryCmd)

	discoveryCmd.PersistentFlags().StringArray("integration-info", []string{}, "Integration info in the form 'integration=AWS,id=123,id_name=name'"+
		"values are optional and support regex")
	discoveryCmd.PersistentFlags().StringArray("resource-type", []string{}, "resource type to discover")
	discoveryCmd.PersistentFlags().Bool("force-full", false, "force to run full discovery. only matters if resource types not defined")
}
