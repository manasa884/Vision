package cmd

type Album struct {
	Name        string
	Path        string
	ParentAlbum string
	SubAlbum    []SubAlbum
	AlbumImages []AlbumImages
}

type SubAlbum struct {
	Name        string
	PathName    string
	AlbumImages []AlbumImages
}

type AlbumImages struct {
	Name string
}

type Folders struct {
	Path       string
	Name       string
	ParentDir  string
	ParentName string
	Remove     bool
}
