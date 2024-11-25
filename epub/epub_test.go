package epub_test

import (
	"reader/epub"
	"testing"
)

func TestLoadPackage(t *testing.T) {
	// <?xml version="1.0" encoding="utf-8" standalone="yes"?>
	data := `
<?xml version="1.0" encoding="UTF-8"?>
<package version="2.0" unique-identifier="pub-id" xmlns="http://www.idpf.org/2007/opf">
  <metadata xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:opf="http://www.idpf.org/2007/opf">
    <dc:identifier id="pub-id" opf:scheme="UUID">urn:uuid:fe93046f-af57-475a-a0cb-a0d4bc99ba6d</dc:identifier>
	<dc:publisher>Your Publisher</dc:publisher>
	<dc:creator opf:file-as="Author, Your" opf:role="aut">Your Author</dc:creator>
    <dc:title>Your title</dc:title>
    <dc:language>en</dc:language>
  </metadata>
  <manifest>
    <item id="ncx" href="toc.ncx" media-type="application/x-dtbncx+xml" />
    <item id="section0001.xhtml" href="xhtml/section0001.xhtml" media-type="application/xhtml+xml" />
	<item id="section0002.xhtml" href="xhtml/section0002.xhtml" media-type="application/xhtml+xml" />
  </manifest>
  <spine toc="ncx">
    <itemref idref="section0001.xhtml" />
	<itemref idref="section0002.xhtml" />
  </spine>
</package>`
	p := epub.LoadContent(data)

	// Metadata
	if p.Metadata.Title != "Your title" {
		t.Errorf("Got: '%s'", p.Metadata.Title)
	}
	if p.Metadata.Author != "Your Author" {
		t.Errorf("Got: '%s'", p.Metadata.Author)
	}
	if p.Metadata.Publisher != "Your Publisher" {
		t.Errorf("Got: '%s'", p.Metadata.Publisher)
	}

	// Manifest
	if p.Manifest.Item[0].Href != "toc.ncx" {
		t.Errorf("Got: '%s'", p.Manifest.Item[0].Href)
	}
	if p.Manifest.Item[0].Id != "ncx" {
		t.Errorf("Got: '%s'", p.Manifest.Item[0].Id)
	}
	if p.Manifest.Item[0].MediaType != "application/x-dtbncx+xml" {
		t.Errorf("Got: '%s'", p.Manifest.Item[0].MediaType)
	}
	if len(p.Manifest.Item) != 3 {
		t.Error("Not enough items in manifest")
	}

	// Spine
	if p.Spine.ItemRef[0].Idref != "section0001.xhtml" {
		t.Errorf("Got: '%s'", p.Spine.ItemRef[0].Idref)
	}
	if len(p.Spine.ItemRef) != 2 {
		t.Error("Not enough items in spine")
	}
}

func TestLoadToc(t *testing.T) {
	data := `
<?xml version="1.0" encoding="UTF-8"?>
<ncx xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1">
  <head>
    <meta name="dtb:uid" content="urn:uuid:fe93046f-af57-475a-a0cb-a0d4bc99ba6d" />
  </head>
  <docTitle>
    <text>Your Title</text>
  </docTitle>
  <navMap>
    <navPoint id="navPoint-1">
      <navLabel>
        <text>Section 1</text>
      </navLabel>
      <content src="xhtml/section0001.xhtml" />
    </navPoint>
	<navPoint id="navPoint-2">
      <navLabel>
        <text>Section 2</text>
      </navLabel>
      <content src="xhtml/section0002.xhtml" />
    </navPoint>
  </navMap>
</ncx>
`
	toc := epub.LoadTableOfContents(data)

	// head
	if toc.Head.Meta[0].Name != "dtb:uid" {
		t.Errorf("Got: '%s'", toc.Head.Meta[0].Name)
	}
	if toc.Head.Meta[0].Content != "urn:uuid:fe93046f-af57-475a-a0cb-a0d4bc99ba6d" {
		t.Errorf("Got: '%s'", toc.Head.Meta[0].Content)
	}

	// Doc Title
	if toc.DocTitle != "Your Title" {
		t.Errorf("Got: '%s'", toc.Head.Meta[0].Content)
	}

	// NavMap
	if toc.NavMap.NavPoint[0].NavLabel[0] != "Section 1" {
		t.Errorf("Got: '%s'", toc.Head.Meta[0].Name)
	}
	if toc.NavMap.NavPoint[0].Content.Src != "xhtml/section0001.xhtml" {
		t.Errorf("Got: '%s'", toc.Head.Meta[0].Content)
	}
	if toc.NavMap.NavPoint[1].NavLabel[0] != "Section 2" {
		t.Errorf("Got: '%s'", toc.Head.Meta[0].Name)
	}
	if toc.NavMap.NavPoint[1].Content.Src != "xhtml/section0002.xhtml" {
		t.Errorf("Got: '%s'", toc.Head.Meta[0].Content)
	}

}

func TestLoadXhtml(t *testing.T) {
	data := `
<?xml version="1.0" encoding="utf-8" standalone="no"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"
  "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">

<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <title>Il Principe</title>
  <link href="../Styles/stylesheet.css" rel="stylesheet" type="text/css" />
  <link href="../Styles/page_styles.css" rel="stylesheet" type="text/css" />
</head>

<body class="calibre">
  <div class="calibre1" id="halftitle_page">
    <h1 class="halftitle" id="calibre_pb_0"><span class="calibre2">BIBLIOTECA IDEALE GIUNTI</span></h1>
  </div>
</body>
</html>
`
	x := epub.LoadXhtml(data)
	if string(x.Head.Content) != `
  <title>Il Principe</title>
  <link href="../Styles/stylesheet.css" rel="stylesheet" type="text/css" />
  <link href="../Styles/page_styles.css" rel="stylesheet" type="text/css" />
` {
		t.Errorf("Got: '%s'", x.Head.Content)
	}

	if string(x.Body.Content) != `
  <div class="calibre1" id="halftitle_page">
    <h1 class="halftitle" id="calibre_pb_0"><span class="calibre2">BIBLIOTECA IDEALE GIUNTI</span></h1>
  </div>
` {
		t.Errorf("Got: '%s'", x.Body.Content)
	}
}
