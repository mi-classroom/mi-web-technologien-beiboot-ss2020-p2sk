package main

const (
	anzahlFarben = 8
	uploadDir    = "uploads/"
	cssDir       = "css/"
	jsDir        = "js/"
	templateDir  = "templates/**/*"
)

var (
	validMimes []string = []string{
		"image/jpeg",
		"image/png",
	}
	defaultMaße []BildMaß = []BildMaß{
		{1024, 0},
		{768, 0},
		{400, 400},
		{360, 0},
	}
	ignoreFiles []string = []string{".gitkeep", "colormap.json"}
)
