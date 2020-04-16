package main

import (
	"image"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

var validMimes []string = []string{
	"image/jpeg",
	"image/png",
}

//type validMimes []string

func isValid(mimeType string) bool {
	for _, mime := range validMimes {
		if mime == mimeType {
			return true
		}
	}
	return false
}

//
type BildTyp struct {
	Typ  string
	Size int
}

// BildTypen
var defaultTypen []BildTyp = []BildTyp{
	BildTyp{"l", 768},
	BildTyp{"m", 640},
	BildTyp{"q", 400},
	BildTyp{"s", 360},
}

// Bild Model
type Bild struct {
	Pfad string
}

func (b *Bild) Dir() string {
	return filepath.Dir(b.Pfad)
}

// Name liefert den Name des Bildes
func (b *Bild) Name() string {
	return filepath.Base(b.Pfad)
}

// Ext liefert die Dateierweiterung des Bildes
func (b *Bild) Ext() string {
	return filepath.Ext(b.Pfad)
}

// MimeType liefert entsprechend der Dateiendung den Mime Typ
func (b *Bild) MimeType() string {
	return mime.TypeByExtension(b.Ext())
}

// Open öffnet die Datei und gibt eine Datei zurück
func (b *Bild) Open() *os.File {
	handler, err := os.Open(b.Pfad)
	if err != nil {
		panic("Datei konnte nicht geöffnet werden")
	}
	return handler
}

//
func (b *Bild) Image() image.Image {
	image, err := imaging.Open(b.Pfad)
	if err != nil {
		panic("Bild konnte nicht geööfnet werden")
	}
	return image
}

//
func (b Bild) Resize(typ BildTyp) {
	resized := imaging.Resize(b.Image(), typ.Size, 0, imaging.Lanczos)
	newFile := filepath.Join(b.Dir(), typ.Typ+"-"+strconv.Itoa(typ.Size)+b.Ext())
	imaging.Save(resized, newFile)
}

//
func (b Bild) Typ() BildTyp {
	dateiName := b.Name()

	for _, item := range defaultTypen {
		if strings.Contains(dateiName, item.Typ) {
			return item
		}
	}
	tmp := strings.Split(strings.Split(dateiName, ".")[0], "-")
	typ := tmp[0]
	size, _ := strconv.Atoi(tmp[1])
	return BildTyp{typ, size}
}

/*func Save(string path) {

}*/
