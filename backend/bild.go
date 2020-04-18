package main

import (
	"image"
	"mime"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
)

var validMimes []string = []string{
	"image/jpeg",
	"image/png",
}

func isValid(mimeType string) bool {
	for _, mime := range validMimes {
		if mime == mimeType {
			return true
		}
	}
	return false
}

// BildMaß
type BildMaß struct {
	Breite int
	Höhe   int
}

func (bm *BildMaß) isQuad() bool {
	return bm.Breite == bm.Höhe
}

func perFaktor(faktor, breite int) BildMaß {
	return BildMaß{breite * faktor / 100, 0}
}

// BildTypen
var defaultMaße []BildMaß = []BildMaß{
	{1024, 0},
	{768, 0},
	{400, 400},
	{360, 0},
}

// Bild Model
type Bild struct {
	Pfad string
}

// Dir liefert den Pfad des Bildes
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

//
func (b *Bild) Width() int {
	config, _, _ := image.DecodeConfig(b.Open())
	return config.Width
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
func (b Bild) Resize(maß BildMaß) {
	resized := imaging.Resize(b.Image(), maß.Breite, maß.Höhe, imaging.Lanczos)
	dx, dy := resized.Bounds().Dx(), resized.Bounds().Dy()

	newFile := filepath.Join(b.Dir(), strconv.Itoa(dx)+"x"+strconv.Itoa(dy)+b.Ext())
	imaging.Save(resized, newFile)
}

//
func (b Bild) CropResize(maß BildMaß) {
	var cropped *image.NRGBA
	image := b.Image()
	dx, dy := image.Bounds().Dx(), image.Bounds().Dy()

	if dx > dy {
		cropped = imaging.CropCenter(image, dy, dy)
	} else {
		cropped = imaging.CropCenter(image, dx, dx)
	}
	resized := imaging.Resize(cropped, maß.Breite, 0, imaging.Lanczos)

	dx, dy = resized.Bounds().Dx(), resized.Bounds().Dy()

	newFile := filepath.Join(b.Dir(), strconv.Itoa(dx)+"x"+strconv.Itoa(dy)+b.Ext())
	imaging.Save(resized, newFile)
}

//
/*func (b Bild) Size() BildTyp {
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
}*/

/*func Save(string path) {

}*/
