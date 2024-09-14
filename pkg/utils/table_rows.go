package utils

import (
	"strconv"

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
			strconv.Itoa(len(benchmark.Metadata.PrimaryTables)),
		}
		rows = append(rows, row)
	}

	return rows

}
