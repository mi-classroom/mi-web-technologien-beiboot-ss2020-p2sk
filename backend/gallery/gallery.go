package gallery

import (
	"image/color"
	"os"
	"path/filepath"
	"time"
)

const dirMode os.FileMode = 0755

// Gallery hält alle Bilder im Uploadverzeichnis
type Gallery map[string]PictureContainer

// PictureContainer stellt Informationen eines Bildes bereit
type PictureContainer struct {
	Dir        string
	ColorMap   ColorPalette
	Collection Collection
}

// CreatePictureFolder ...
func CreatePictureFolder(uploadDir string) string {
	dirName := time.Now().Format("20060102150405")
	path := filepath.Join(uploadDir, dirName)
	os.Mkdir(path, dirMode)
	return path
}

// ToScrset liefert pro Sammlung den scrset
/*func (pc PictureContainer) ToScrset() string {
	var scrset []string
	for _, i := range pc.Collection {
		scrset = append(scrset, i.Path+" "+strconv.Itoa(i.Width())+"w")
	}
	return strings.Join(scrset, ", ")
}*/

// Collection stellt zu einem Bild alle Referenzen/Größen dar
type Collection []Picture

// GetPreviewPicture ...
func (c Collection) GetPreviewPicture() Picture {
	for _, image := range c {
		if image.Size().IsQuad() {
			return image
		}
	}
	return c[0]
}

// ColorPalette beschreibt eine Sammlung von Farben
type ColorPalette []color.RGBA

// NewColorPalette ...
func NewColorPalette(palette color.Palette) ColorPalette {
	fp := make(ColorPalette, len(palette))
	for i, c := range palette {
		fp[i] = c.(color.RGBA)
	}
	return fp
}
