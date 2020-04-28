package main

const (
	anzahlFarben = 8
	uploadDir    = "uploads/"
	templateDir  = "templates/**/*"
	colorFile    = "colormap.json"
)

var (
	validMimes []MimeType = []MimeType{
		"image/jpeg",
		"image/png",
	}
	defaultMaße []BildMaß = []BildMaß{
		{1024, 0},
		{768, 0},
		{400, 400},
		{360, 0},
	}
	ignoreFiles []string = []string{".gitkeep"}
)
