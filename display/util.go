package display

import (
	"log"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func ConvertHtmlToMarkdown(html string) string {
	markdown, err := htmltomarkdown.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	return markdown
}
