package display

import (
	// "fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	// "github.com/charmbracelet/lipgloss"
	"log"
	"reader/epub"
)

type ReadingModel struct {
	Opened   bool
	Book     epub.Epub
	Current  string
	viewport viewport.Model
}

func NewReadingModel(width int, height int) ReadingModel {
	// book := epub.LoadBook("data/linear-algebra.epub")
	book := epub.LoadBook("data/Il-Principe.epub")
	first := ""
	for key := range book.Content {
		first = key
		break
	}

	vp := viewport.New(width, height)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width-2),
	)
	if err != nil {
		log.Fatal(err)
	}

	content := book.RenderContentHtml(first)
	content = ConvertHtmlToMarkdown(content)
	str, err := renderer.Render(content)
	if err != nil {
		log.Fatal(err)
	}

	vp.SetContent(str)

	return ReadingModel{
		Opened:   true,
		Book:     book,
		Current:  first,
		viewport: vp,
	}

}

func (m ReadingModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m ReadingModel) Update(msg tea.Msg) (ReadingModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m ReadingModel) View() string {
	return m.viewport.View() //+ m.helpView()
}
