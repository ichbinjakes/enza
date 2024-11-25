package epub

import (
	"encoding/xml"
	"log"
	"fmt"
)
// https://idpf.org/epub/20/spec/OPS_2.0.1_draft.htm

// OPS Content

// https://www.w3.org/TR/xhtml11/doctype.html#s_doctype
type XHTML struct {
	XMLName xml.Name `xml:"html"`
	// Inner string `xml:",innerxml"`
	Head XHTMLHead `xml:"head"`
	Body XHTMLBody `xml:"body"`
}

type XHTMLHead struct {
	XMLName xml.Name `xml:"head"`
	Content string `xml:",innerxml"`
}

type XHTMLBody struct {
	XMLName xml.Name `xml:"body"`
	Content string `xml:",innerxml"`
}

// type HTML struct {
// 	Head []byte `xml:"head"`
// 	Body []by
// }
// https://groups.niso.org/higherlogic/ws/public/download/14650/Z39_86_2005r2012.pdf
type DTBook struct {
}

// OPS Publication

type OPS struct {
	Package string
	Toc     string
	Content []string
}


func LoadXhtml(content string) XHTML {
	x := XHTML{}
	err := xml.Unmarshal([]byte(content), &x)
	if err != nil {
		log.Fatal(err)
	}
	// println("INNER:", string(x.Inner))
	// println("HEAD:", string(x.Head.Content))
	// println("BODY:", string(x.Body.Content))
	// println("HTML:", string(x.Html))
	return x
}

func (x XHTML) GetHTML() string {
	output, err := xml.Marshal(x)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return string(output)
}