package epub

import (
	"archive/zip"
	// "encoding/xml"
	"errors"
	"io"
	"log"
	"regexp"
	"path"
	// "encoding/xml"
)


type OCFContainer map[string]string

type Epub struct {
	MetaInf MetaInf
	Package Package
	Toc     TableOfContent
	Content map[string]string
	OCF OCFContainer
}

type MetaInf struct {
	Container Container
}

type Container struct {
	RootFile string
}

// Load an epub into an OCR container instance
func LoadOCFContainer(path string) OCFContainer {

	// open the archive
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	epub := make(OCFContainer)
	for _, f := range r.File {
		// Skip adding directories to epub
		if f.FileInfo().IsDir() {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		b, err := io.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}

		epub[f.Name] = string(b)
	}

	return epub
}

func (ocf OCFContainer) GetPackagePath() (string, error) {
	// get the package manifest
	re := regexp.MustCompile(`(.)*.opf`)
	for key := range ocf {
		match := re.MatchString(key)
		if match {
			return key, nil
		}
	}
	return "", errors.New("ocf package path not found")
}

func (ocf OCFContainer) GetTableOfContentsPath() (string, error) {
	// Search for toc.ncx
	re := regexp.MustCompile(`(.)*/toc.ncx`)
	for key := range ocf {
		match := re.MatchString(key)
		if match {
			return key, nil
		}
	}

	// Try get TOC from package file
	packagePath, err := ocf.GetPackagePath()
	if err == nil {
		p := LoadPackageManifest(ocf[packagePath])
		
		// Try get TOC from spine
		tocManifestId := p.Spine.TOC
		if tocManifestId != "" {
			for _, value := range p.Manifest.Item {
				if value.Id == tocManifestId {
					return value.Href, nil
				}
			}
		}

		// Try and get TOC from manifest
		for _, value := range p.Manifest.Item {
			if value.MediaType == "application/x-dtbncx+xml" {
				return value.Href, nil
			}
		}
	}

	return "", errors.New("table of contents path not found")
}

func (ocf OCFContainer) GetContentDirectory() string {
	packagePath, err := ocf.GetPackagePath()
	if err != nil {
		log.Fatal("failed to find the package path")
	}
	return path.Dir(packagePath)

}

func LoadBook(bookPath string) Epub {
	ctr := LoadOCFContainer(bookPath)

	book := Epub{
		OCF: ctr,
		Content: make(map[string]string),
	}

	packagePath, err := ctr.GetPackagePath()
	if err != nil {
		log.Fatal("failed to find the package path")
	}
	book.Package = LoadPackageManifest(ctr[packagePath])

	// we don't do anything with this?
	tocPath, err := ctr.GetTableOfContentsPath()
	if err == nil {
		book.Toc = LoadTableOfContents(ctr[tocPath])
	}

	contentDir := ctr.GetContentDirectory()
	for _, value := range book.Package.Manifest.Item {
		book.Content[path.Join(contentDir, value.Href)] = ctr[path.Join(contentDir, value.Href)]
	}
	
	return book
}


func (book Epub) RenderContentHtml(path string) string {
	raw, ok := book.Content[path]
	if !ok {
		println(path)
		log.Fatal("could not find path")
	}
	content := LoadXhtml(raw)
	return content.GetHTML()
	// err := xml.Unmarshal([]byte(raw), &content)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// println("HTML", content.Html)
	// return string(content.Html)
}