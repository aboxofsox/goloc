package goloc

import (
	"log"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyName    = "Extension"
	columnKeyElement = "LoC"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#fff")).
			Background(lipgloss.Color("#ff0084")).
			PaddingLeft(1).
			PaddingRight(1).
			MarginBottom(1).
			MarginTop(1)

	totalStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ff0084")).
			MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666")).
			Italic(true)

	total int = 0
)

type Model struct {
	table table.Model
}

func NewTable(m map[string]int) Model {
	sorted := sorter(m)
	cols := []table.Column{
		table.NewColumn(columnKeyName, "Extension", 12).WithStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#4293f5"))),
		table.NewColumn(columnKeyElement, "LoC", 10).WithStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#ffff00"))),
	}
	rows := []table.Row{}

	for _, k := range sorted {
		total += m[k]
		rows = append(rows, table.NewRow(table.RowData{
			columnKeyName:    k,
			columnKeyElement: m[k],
		}).WithStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#fff"))))
	}
	tbl := table.New(cols).WithRows(rows)
	return Model{
		table: tbl,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString(titleStyle.Render("Goloc"))
	body.WriteString("\n")
	body.WriteString(m.table.View())
	body.WriteString("\n")
	body.WriteString(totalStyle.Render("Total: ") + strconv.FormatInt(int64(total), 10))
	body.WriteString(infoStyle.Render("\nPress q to quit\n"))

	return body.String()
}

func MakeTable(m map[string]int) {
	p := tea.NewProgram(NewTable(m))

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

}
