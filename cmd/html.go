package cmd

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
)

//HTML template stuff goes here.

func CreateIndex(album Album) {
	box := packr.NewBox("../template")
	// templateThing := filepath.Join(templatePath, "index.tmpl")
	// s, err := box.FindString("index.tmpl")
	// if err != nil {
	// 	log.Fatal(err)
	// }

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

}

// func DefaultTemple(path string) {

// 	indexFile := filepath.Join(path, "index.tmpl")
// 	f, _ := os.Create(indexFile)
// 	f.Write([]byte("FART: {{.Name}}\n{{ range .SubAlbum}}{{$subName := .Name}}\n<div class=\"codrops-top clearfix\">\n<div class=\"albums-tab\">\n<a href=\"./{{.Name}}\">\n<div class=\"albums-tab-thumb sim-anim-7\">\n{{range .AlbumImages}}\n<img src=\"./{{$subName}}/visionimg/thumb_{{.Name}}\" class=\"all studio\" />\n{{end}}\n</div>\n<div class=\"albums-tab-text\">{{.Name}}</div>\n</a>\n</div>\n</div>{{end}}"))
// 	f.Close()
// }
