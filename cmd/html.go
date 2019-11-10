package cmd

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
)

//HTML template stuff goes here.

//CreateIndex pulls in the template to the package.
func CreateIndex(album Album) {
	box := packr.NewBox("../template")

	t, err := template.New("index").Parse(box.String("index.tmpl"))

	if err != nil {
		log.Print(err)
		return
	}
	indexFile := filepath.Join(album.Path, "index.html")
	f, err := os.Create(indexFile)
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = t.Execute(f, album)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
	cssFile := []byte(box.String("styles.css"))
	if err != nil {
		log.Print(err)
		return
	}
	cssFilePath := filepath.Join(album.Path, "visionimg/styles.css")
	err = ioutil.WriteFile(cssFilePath, cssFile, 0777)
	if err != nil {
		log.Print(err)
		return
	}
}
