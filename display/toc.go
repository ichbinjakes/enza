package display

import (
	// "fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbles/viewport"
	// "github.com/charmbracelet/glamour"
	// "log"
	// "reader/epub"
)

type TOCModel struct {
	list     list.Model
	index    int
	selected int
}

func NewTOCModel() TOCModel {
	items := []list.Item{
		TOCItem{title: "hello", path: "a"},
		TOCItem{title: "world", path: "b"},
	}
	return TOCModel{
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		index:    0,
		selected: 0,
	}
}

type TOCItem struct {
	title string
	path  string
}

func (t TOCItem) FilterValue() string {
	return t.title
}

func (m TOCModel) Init() tea.Cmd {
	return nil
}

func (m TOCModel) Update(msg tea.Msg) (TOCModel, tea.Cmd) {
	return m, nil
}

func (m TOCModel) View() string {
	return m.list.View()
}
