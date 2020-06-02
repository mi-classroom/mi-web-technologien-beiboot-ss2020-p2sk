package config

import "../gallery"

const (
	ColorCount  = 8
	UploadDir   = "uploads/"
	TemplateDir = "templates/**/*"
	ColorFile   = "colormap.json"
)

var (
	DefaultImageSizes []gallery.ImageSize = []gallery.ImageSize{
		{1024, 0},
		{768, 0},
		{400, 400},
		{360, 0},
	}
	IgnoreFiles []string = []string{".gitkeep"}
)
