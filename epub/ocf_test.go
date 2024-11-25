package epub_test

import (
	"reader/epub"
	"testing"
	"fmt"
)

// func TestLoadEpub(t *testing.T) {
// 	file := epub.LoadEpub("../data/linear-algebra.epub")
// 	if len(file) != 99 {
// 		t.Error("Not enough items loaded from linear algebra epub.")
// 	}
// }

func TestLoadPackagePath(t *testing.T) {
	ctr := epub.LoadOCFContainer("../data/linear-algebra.epub")
	path, err := ctr.GetPackagePath()
	if err != nil {
		fmt.Println(err)
		t.Error("could not find linear algebra epub's package path.")
	}	
	if path != "linear-algebra/EPUB/package.opf" {
		fmt.Println(path)
		t.Error("could not find linear algebra epub's correct package path.")
	}
}

func TestLoadBook(t *testing.T) {
	file := epub.LoadBook("../data/linear-algebra.epub")
	if len(file.OCF) != 99 {
		t.Error("Not enough items loaded from linear algebra epub.")
	}
}

