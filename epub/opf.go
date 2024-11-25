package epub

// https://idpf.org/epub/20/spec/OPF_2.0.1_draft.htm

import (
	"log"
	// "regexp"
	"encoding/xml"
	// "fmt"
)

// OPF Package Document

type Package struct {
	XMLName  xml.Name `xml:"http://www.idpf.org/2007/opf package"`
	Metadata Metadata `xml:"metadata"`
	Manifest Manifest `xml:"manifest"`
	Spine    Spine    `xml:"spine"`
	Guide    Guide    `xml:"guide"`
}

// type Title struct {
// 	XMLName xml.Name `xml:"dc:title"`
// 	Title string
// }

type Metadata struct {
	Inner     []byte `xml:",innerxml"`
	Title     string `xml:"http://purl.org/dc/elements/1.1/ title"`
	Author    string `xml:"http://purl.org/dc/elements/1.1/ creator"`
	Publisher string `xml:"http://purl.org/dc/elements/1.1/ publisher"`
}

type ManifestItem struct {
	Href      string `xml:"href,attr"`
	Id        string `xml:"id,attr"`
	MediaType string `xml:"media-type,attr"`
}

type Manifest struct {
	Item []ManifestItem `xml:"item"`
}

type SpineItemref struct {
	Idref string `xml:"idref,attr"`
}

type Spine struct {
	ItemRef []SpineItemref `xml:"itemref"`
	TOC     string         `xml:"toc,attr"`
}

type GuideReference struct {
	Href  string `xml:"href,attr"`
	Title string `xml:"title,attr"`
	Type  string `xml:"type,attr"`
}

type Guide struct {
	Reference []GuideReference `xml:"reference"`
}

// TOC

type Head struct {
	Meta []Meta `xml:"meta"`
}

type Meta struct {
	Inner   []byte `xml:",innerxml"`
	Name    string `xml:"name,attr"`
	Content string `xml:"content,attr"`
}

type NavMap struct {
	NavPoint []NavPoint `xml:"navPoint"`
	// NavInfo  []string  `xml:"navInfo>text"`
	// NavLabel []string `xml:"navLabel>text"`
}

type NavInfo struct {
	Text string
}

type NavPoint struct {
	NavLabel []string `xml:"navLabel>text"`
	Content  Content  `xml:"content"`
	// NavPoint  []NavPoint `xml:"navPoint"`
	Class     string `xml:"class,attr"`
	Id        string `xml:"id,attr"`
	PlayOrder string `xml:"playOrder,attr"`
}

type NavLabel struct {
	Text string
}

type Content struct {
	Id  string `xml:"id,attr"`
	Src string `xml:"src,attr"`
}

type TableOfContent struct {
	XMLName   xml.Name `xml:"http://www.daisy.org/z3986/2005/ncx/ ncx"`
	Head      Head     `xml:"head"`
	DocTitle  string   `xml:"docTitle>text"`
	DocAuthor string   `xml:"docAuthor>text"`
	NavMap    NavMap   `xml:"navMap"`
}

// Utilities
func LoadPackageManifest(content string) Package {
	p := Package{}
	err := xml.Unmarshal([]byte(content), &p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func LoadTableOfContents(content string) TableOfContent {
	t := TableOfContent{}
	err := xml.Unmarshal([]byte(content), &t)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func LoadContent(content string) Package {
	p := Package{}
	err := xml.Unmarshal([]byte(content), &p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
