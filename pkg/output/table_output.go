package output

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
	tableColor   = lipgloss.Color("165")
	oddRowColor  = lipgloss.Color("254")
	evenRowColor = lipgloss.Color("245")
)

func PrintControlsTable(rows [][]string) {

	re := lipgloss.NewRenderer(os.Stdout)

	var (
		HeaderStyle  = re.NewStyle().Foreground(tableColor).Bold(true).Align(lipgloss.Center)
		CellStyle    = re.NewStyle().Padding(0, 1).Width(30)
		OddRowStyle  = CellStyle.Foreground(oddRowColor)
		EvenRowStyle = CellStyle.Foreground(evenRowColor)
	)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		}).
		Headers("CONTROL ID", "CONTROL TABLE", "SEVERITY", "PRIMARY TABLE").
		Rows(rows...)

	fmt.Println(t)

}
