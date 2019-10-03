package cmd

//Album information
type Album struct {
	Name        string
	Path        string
	ParentAlbum string
	SubAlbum    []SubAlbum
	AlbumImages []AlbumImages
}

//SubAlbum information
type SubAlbum struct {
	Name        string
	PathName    string
	AlbumImages []AlbumImages
}

//AlbumImages name
type AlbumImages struct {
	Name string
}

//Folders information
type Folders struct {
	Path       string
	Name       string
	ParentDir  string
	ParentName string
	Remove     bool
}
