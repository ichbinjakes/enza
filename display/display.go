package display

import (
	// "fmt"
	"log"
	"os"

	// "reader/epub"
	// "github.com/charmbracelet/bubbles/viewport"
	// "github.com/charmbracelet/glamour"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model
// 	INIT - a function that returns initial command for the application to run
// 	UPDATE - a function that handles incoming events and updates model accordingly, returns the model and a cmd
// 	VIEW - a function that renders the UI based on the data in the model
// Cmds perform I/O and return a msg

var (
	tocWidth       = 40
	tocHeight      = 100
	
	readingWidth       = 100
	readingHeight      = 50
	
)

type MainModel struct {
	TOC          TOCModel
	ReadingModel ReadingModel
	activeModel  int // 0 -> toc, 1 - > reader
}

func (m MainModel) dispatchUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.activeModel == 0 {
		_m, cmd := m.TOC.Update(msg)
		m.TOC = _m
		return m, cmd
	}
	if m.activeModel == 1 {
		_m, cmd := m.ReadingModel.Update(msg)
		m.ReadingModel = _m
		return m, cmd
	}
	println("HECK")
	log.Fatal("cannot find active pane")
	os.Exit(1)
	return nil, nil
}

func NewMainModel() MainModel {
	return MainModel{
		TOC:          NewTOCModel(),
		ReadingModel: NewReadingModel(readingWidth, readingHeight),
		activeModel:  0,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "left":
			if m.activeModel == 1 {
				m.activeModel = 0
				return m, nil
			}
			return m, nil
		case "right":
			if m.activeModel == 0 {
				m.activeModel = 1
				return m, nil
			}
			return m, nil
		default:
			return m.dispatchUpdate(msg)
		}
	default:
		return m, nil
	}
}

func (m MainModel) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.getTocStyle().Render(m.TOC.View()),
		m.getReaderStyle().Render(m.ReadingModel.View()),
	)
}

func (m MainModel) getTocStyle() lipgloss.Style {
	if m.activeModel == 0 {
		return lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("4")).
			PaddingRight(2)
	} else {
		return lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("8")).
			PaddingRight(2)
	}
}

func (m MainModel) getReaderStyle() lipgloss.Style {
	if m.activeModel == 1 {
		return lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("4")).
			PaddingRight(6).
			PaddingLeft(6).
			PaddingTop(2).
			PaddingBottom(2)
	} else {
		return lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("8")).
			PaddingRight(6).
			PaddingLeft(6).
			PaddingTop(2).
			PaddingBottom(2)
	}
}
