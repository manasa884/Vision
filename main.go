package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/LordBrain/Vision/cmd"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("vision", "Tool to create static image albums")

	create     = app.Command("create", "Create a new album")
	createPath = create.Arg("path", "Path to images").Required().ExistingDir()
	// createTemplate = create.Flag("template", "Template to use").String()
	createWidth = create.Flag("width", "Resized image width").Short('w').Int()

	update         = app.Command("update", "Update existing album")
	updatePath     = update.Arg("path", "Path to album").Required().ExistingDir()
	updateTemplate = update.Flag("template", "Template to use").String()
)

func main() {
	// kingpin.Parse()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Create new album
	case create.FullCommand():
		println("Creating")
		path, _ := filepath.Abs(*createPath)
		startPath := path
		fmt.Println("Start Path:", startPath)
		folders := cmd.GetFolders(path)

		allAlbums := cmd.GenAlbums(startPath, folders)
		fmt.Println("Number of album entries:", len(allAlbums))

		for _, album := range allAlbums {
			fmt.Printf("Album: %s\nNumber of images: %d\n", album.Name, len(album.AlbumImages))
			fmt.Println("Album Path:", album.Path)
			if len(album.SubAlbum) > 0 {
				for _, subalbum := range album.SubAlbum {
					fmt.Println("Sub album name:", subalbum.Name)
				}
			}
			//Resize images and create thumbnails
			for _, imageName := range album.AlbumImages {
				imagePath := filepath.Join(album.Path, imageName.Name)
				if strconv.Itoa(*createWidth) != "0" {
					cmd.GenImages(imagePath, *createWidth)
				} else {
					cmd.GenImages(imagePath, 800)
				}

			}

			//Create html files

			cmd.CreateIndex(album)

			fmt.Println("----")
		}

	// Update album
	case update.FullCommand():
		println("Updating")
	}
}
