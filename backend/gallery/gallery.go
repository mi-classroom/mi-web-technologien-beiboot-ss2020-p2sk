package gallery

import (
	"encoding/json"
	"image/color"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/RobCherry/vibrant"
)

const dirMode os.FileMode = 0755

const (
	ALPHA  SortType = "alpha"
	COLOR           = "color"
	DATE            = "date"
	RANDOM          = "random"
)

type SortType string

// Gallery hält alle Bilder im Uploadverzeichnis
//type Gallery map[string]PictureContainer
type Gallery []PictureContainer

func (g Gallery) Sort(by SortType) {
	switch by {
	case ALPHA:
		sort.Slice(g, func(i, j int) bool {
			return g[i].Dir > g[j].Dir
		})

	case COLOR:
		sort.Slice(g, func(i, j int) bool {
			return g[i].ColorMap.Vibrant().HSL.H > g[j].ColorMap.Vibrant().HSL.H
		})
	case DATE:
		sort.Slice(g, func(i, j int) bool {
			return g[i].Dir < g[j].Dir
		})
	case RANDOM:
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(g), func(i, j int) { g[i], g[j] = g[j], g[i] })
	default:
	}
}

func (g Gallery) Reduce(from, count int) Gallery {
	if count > len(g) {
		return g
	}
	if from > len(g) {
		return g[0:count]
	}

	to := count + from
	if to > len(g) {
		return g[len(g)-count:]
	}
	return g[from:to]
}

// LoadGallery lädt alle Bilder und liefert ein Gallery Objekt zurück
func LoadGallery(imageBaseDir string, colorFile string, ignoreFiles []string) Gallery {
	var currentDir string
	var tempContainer PictureContainer
	tempGallery := make(Gallery, 0)

	filepath.Walk(imageBaseDir, func(path string, info os.FileInfo, err error) error {
		// das Uploadverzeichnis ignorieren
		if info.IsDir() && info.Name() == filepath.Dir(imageBaseDir) {
			return err
		}

		// weitere Dateien ignorieren
		if stringInSlice(ignoreFiles, filepath.Base(path)) {
			return nil
		}

		// wurde ein Verzeichnis gefunden neues PictureContainer erstellen
		if info.IsDir() {
			if currentDir != "" {
				tempGallery = append(tempGallery, tempContainer)
			}
			//tempGallery[lastDir]
			currentDir = info.Name()
			tempContainer = PictureContainer{
				Dir:        currentDir,
				ColorMap:   ColorPalette{},
				Collection: make([]Picture, 0),
			}
			return nil
		}

		if info.Name() == colorFile {
			var rgba ColorPalette
			//tmp := tempGallery[lastDir]
			data, _ := ioutil.ReadFile(path)

			json.Unmarshal(data, &rgba)
			//tmp.ColorMap = rgba
			tempContainer.ColorMap = rgba

			return nil
		}

		// Ab hier haben wir ein Bild gefunden
		// Pfadseperator / Unix/Windows
		tempContainer.Collection = append(tempContainer.Collection, Picture{Path: filepath.ToSlash(path)})
		return nil
	})

	tempGallery = append(tempGallery, tempContainer)
	return tempGallery
}

func stringInSlice(s []string, needle string) bool {
	for _, item := range s {
		if item == needle {
			return true
		}
	}
	return false
}

// PictureContainer stellt Informationen eines Bildes bereit
type PictureContainer struct {
	Dir        string       `json:"id"`
	ColorMap   ColorPalette `json:"colors"`
	Collection Collection   `json:"images"`
}

// CreatePictureFolder ...
func CreatePictureFolder(uploadDir string) string {
	dirName := time.Now().Format("20060102150405")
	path := filepath.Join(uploadDir, dirName)
	os.Mkdir(path, dirMode)
	return path
}

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

// Color repräsentiert eine quantifizierte Farbe
type Color struct {
	RGBA     color.NRGBA `json:"rgba"`
	HSL      vibrant.HSL `json:"hsl"`
	Quantity uint32      `json:"quantity"`
	Vibrant  string      `json:"vibrant"`
}

// ColorPalette beschreibt die häufigsten Farben des Bildes in unterschiedlichen Formaten
type ColorPalette []Color

// Vibrant gibt die Vibrant Farbe, oder wenn diese nicht vorhanden ist die häufigste zurück.
func (cp ColorPalette) Vibrant() Color {
	var vibrant Color
	var mostCommon Color

	for _, color := range cp {
		if color.Vibrant == "vibrant" {
			vibrant = color
		}

		if mostCommon.Quantity < color.Quantity {
			mostCommon = color
		}
	}

	if &vibrant == nil {
		return mostCommon
	}

	return vibrant
}

// NewColorPalette erzeugt ein neues ColorPalette Objekt
/* func NewColorPalette(palette color.Palette) ColorPalette {
	fp := make(ColorPalette, len(palette))
	for i, c := range palette {
		fp[i] = c.(color.NRGBA)
	}
	return fp
} */
