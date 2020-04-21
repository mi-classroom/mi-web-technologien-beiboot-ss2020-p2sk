package main

import (
	"image"
	"image/color"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/ericpauley/go-quantize/quantize"
)

// Galerie hält alle Bilder im Uploadverzeichnis
type Galerie struct {
	Dir      string
	Sammlung Sammlung
}

// ToScrset liefert pro Sammlung den scrset
func (c Galerie) ToScrset() string {
	var scrset []string
	for _, e := range c.Sammlung {
		scrset = append(scrset, e.Pfad+" "+strconv.Itoa(e.Width())+"w")
	}
	return strings.Join(scrset, ", ")
}

// Sammlung beschreibt die Menge eines Bildes
type Sammlung []Bild

//
func isValid(mimeType string) bool {
	for _, mime := range validMimes {
		if mime == mimeType {
			return true
		}
	}
	return false
}

// Farbpalette beschreibt eine Sammlung von Farben (interface type Color)
type Farbpalette color.Palette

func (f Farbpalette) convert() []struct{ r, g, b, a uint8 } {
	var tmp []struct{ r, g, b, a uint8 }

	for _, c := range f {
		r, g, b, a := c.RGBA()
		r, g, b, a = r>>8, g>>8, b>>8, a>>8
		tmp = append(tmp, struct{ r, g, b, a uint8 }{uint8(r), uint8(g), uint8(b), uint8(a)})
	}
	log.Print(tmp)
	return tmp
}

// BildMaß repräsentiert ein BildMaß aus Breite und Höhe
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

// Width gibt die Breite des Bildes per image.DecodeConfig zurück.
// Diese Funktion ist wesentlich performanter als die Breite über
// den image.Image Typ und desen Bounds() Methode zu holen
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

// Image gibt eine image.Image Repräsentation des Bildes zurück
func (b *Bild) Image() image.Image {
	image, err := imaging.Open(b.Pfad)
	if err != nil {
		panic("Bild konnte nicht geööfnet werden")
	}
	return image
}

// Resize skaliert das Bild auf das angegebene BildMaß
func (b Bild) Resize(maß BildMaß) {
	resized := imaging.Resize(b.Image(), maß.Breite, maß.Höhe, imaging.Lanczos)
	dx, dy := resized.Bounds().Dx(), resized.Bounds().Dy()

	newFile := filepath.Join(b.Dir(), strconv.Itoa(dx)+"x"+strconv.Itoa(dy)+b.Ext())
	imaging.Save(resized, newFile)
}

// CropResize schneidet das Bild quadratisch zu und
// skaliert das Bild auf das angegebene BildMaß
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
func (b Bild) Quantisiere() Farbpalette {
	q := quantize.MedianCutQuantizer{}
	return Farbpalette(q.Quantize(make([]color.Color, 0, anzahlFarben), b.Image()))
}
