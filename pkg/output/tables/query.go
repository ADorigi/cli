package tables

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func PrintQueryResultTable(headers []string, rows [][]string) {

	re := lipgloss.NewRenderer(os.Stdout)

	var (
		HeaderStyle  = re.NewStyle().Foreground(tableColor).Bold(true).Align(lipgloss.Center)
		CellStyle    = re.NewStyle().Padding(0, 1)
		OddRowStyle  = CellStyle.Foreground(oddRowColor)
		EvenRowStyle = CellStyle.Foreground(evenRowColor)
	)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style
			switch {
			case row == 0:
				style = HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle
			default:
				style = OddRowStyle
			}

			return style
		}).
		Headers(headers...).
		Rows(rows...)

	fmt.Println(t)

}
