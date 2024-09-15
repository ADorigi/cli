package tables

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

			switch {
			case col == 0:
				style.Width(40)
			case col == 1:
				style.Width(40)
			case col == 2:
				style.Width(10)
			case col == 3:
				style.Width(40)
			}
			return style
		}).
		Headers("CONTROL ID", "CONTROL TABLE", "SEVERITY", "PRIMARY TABLE").
		Rows(rows...)

	fmt.Println(t)

}

// Benchmark ID, Benchmark Title, # of Controls, # of Primary Tables, Connectors
func PrintBenchmarksTable(rows [][]string) {

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

			switch {
			case col == 0:
				style.Width(40)
			case col == 1:
				style.Width(40)
			case col == 2:
				style.Width(10)
			case col == 3:
				style.Width(10)
			case col == 4:
				style.Width(10)
			}
			return style
		}).
		Headers("BENCHMARK ID", "BENCHMARK TITLE", "# OF CONTROLS", "# OF PRIMARY TABLES", "CONNECTORS").
		Rows(rows...)

	fmt.Println(t)
}
