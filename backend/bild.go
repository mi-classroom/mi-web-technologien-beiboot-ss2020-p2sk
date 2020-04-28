package main

import (
	"encoding/json"
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

// @todo Saubere Typenstruktur erstellen

//
type MimeType string

//
func isValid(mimeType MimeType) bool {
	for _, mime := range validMimes {
		if mime == mimeType {
			return true
		}
	}
	return false
}

// Galerie hält alle Bilder im Uploadverzeichnis
type Galerie struct {
	Dir      string
	ColorMap Farbpalette
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

// Sammlung stellt zu einem Bild alle Referenzen/Größen dar
type Sammlung []Bild

// Farbpalette beschreibt eine Sammlung von Farben
type Farbpalette []color.RGBA

func NewFarbpalette(palette color.Palette) Farbpalette {
	fp := make(Farbpalette, len(palette))
	for i, c := range palette {
		fp[i] = c.(color.RGBA)
	}
	return fp
}

/*func (f Farbpalette) UnmarshalJSON(data []byte) error {
	//var rgba []color.RGBA
	json.Unmarshal(data, f)
	return nil
}*/

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

// Quantisiere erstellt eine Farbpalette
func (b Bild) Quantisiere(anzahl int) Farbpalette {
	q := quantize.MedianCutQuantizer{}
	p := q.Quantize(make([]color.Color, 0, anzahl), b.Image())
	fp := NewFarbpalette(p)
	return fp
}

//
func (b Bild) SpeicherColorMap(palette Farbpalette, fileName string) {
	byte, err := json.Marshal(palette)

	if err != nil {
		log.Panicf("Json Marshal fehlgeschlagen: %s", err.Error())
	}

	jsonFile, _ := os.Create(filepath.Join(b.Dir(), fileName))
	jsonFile.Write(byte)
	jsonFile.Close()
}
