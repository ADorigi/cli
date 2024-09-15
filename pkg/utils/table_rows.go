package utils

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/adorigi/opengovernance/pkg/types"
)

func GenerateControlRows(controls []types.Control) [][]string {

	var rows [][]string

	for _, control := range controls {
		row := []string{
			control.ID,
			control.Title,
			control.Severity,
			control.Query.PrimaryTable,
		}
		rows = append(rows, row)
	}

	return rows

}

func GenerateBenchmarkRows(benchmarks []types.BenchMark) [][]string {

	var rows [][]string

	for _, benchmark := range benchmarks {
		row := []string{
			benchmark.Metadata.ID,
			benchmark.Metadata.Title,
			strconv.Itoa(benchmark.Metadata.NumberOfControls),
			strconv.Itoa(len(benchmark.Metadata.PrimaryTables)),
			strings.Join(benchmark.Metadata.Connectors, ","),
		}
		rows = append(rows, row)
	}

	return rows

}

func GenerateDiscoveryJobsRows(jobs []types.RunDiscoveryResponse) ([][]string, error) {
	var rows [][]string
	for _, job := range jobs {
		integrationInfo, err := json.Marshal(job.IntegrationInfo)
		if err != nil {
			return nil, err
		}
		row := []string{
			strconv.Itoa(int(job.JobId)),
			job.ResourceType,
			job.Status,
			string(integrationInfo),
			job.FailureReason,
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func GenerateComplianceJobsRows(job types.RunBenchmarkResponse) ([][]string, error) {
	var rows [][]string
	integrationInfo, err := json.Marshal(job.IntegrationInfo)
	if err != nil {
		return nil, err
	}
	row := []string{
		strconv.Itoa(int(job.JobId)),
		job.BenchmarkId,
		string(integrationInfo),
	}
	rows = append(rows, row)

	return rows, nil
}
