package main

import "./gallery"

const (
	colorCount  = 8
	uploadDir   = "uploads/"
	templateDir = "templates/**/*"
	colorFile   = "colormap.json"
)

var (
	defaultImageSizes []gallery.ImageSize = []gallery.ImageSize{
		{1024, 0},
		{768, 0},
		{400, 400},
		{360, 0},
	}
	ignoreFiles []string = []string{".gitkeep"}
)
