package utils

import "github.com/adorigi/opengovernance/pkg/types"

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
