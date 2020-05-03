package gallery

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

// MimeType beschreibt das Konzept des MimeType
type MimeType string

// ValidMimes fasst die gültigen Imagem Mimes zusammen
var ValidMimes []MimeType = []MimeType{
	"image/jpeg",
	"image/png",
}

// Hinterfragen, wie zu erweitern?
func IsValid(mimeType MimeType) bool {
	for _, mime := range ValidMimes {
		if mime == mimeType {
			return true
		}
	}
	return false
}

// Gallery hält alle Bilder im Uploadverzeichnis
type Gallery struct {
	Dir        string
	ColorMap   ColorPalette
	Collection Collection
}

// ToScrset liefert pro Sammlung den scrset
func (g Gallery) ToScrset() string {
	var scrset []string
	for _, i := range g.Collection {
		scrset = append(scrset, i.Path+" "+strconv.Itoa(i.Width())+"w")
	}
	return strings.Join(scrset, ", ")
}

// Collection stellt zu einem Bild alle Referenzen/Größen dar
type Collection []Image

// ColorPalette beschreibt eine Sammlung von Farben
type ColorPalette []color.RGBA

// NewColorPalette
func NewColorPalette(palette color.Palette) ColorPalette {
	fp := make(ColorPalette, len(palette))
	for i, c := range palette {
		fp[i] = c.(color.RGBA)
	}
	return fp
}

// ImageSize repräsentiert ein Bildmaß aus Breite und Höhe
type ImageSize struct {
	Width  int
	Height int
}

// IsQuad
func (is *ImageSize) IsQuad() bool {
	return is.Width == is.Height
}

// FromFactor
func FromFactor(factor, width int) ImageSize {
	return ImageSize{factor * width / 100, 0}
}

// Image Model
type Image struct {
	Path string
}

// Dir liefert den Pfad des Bildes
func (i *Image) Dir() string {
	return filepath.Dir(i.Path)
}

// Name liefert den Name des Bildes
func (i *Image) Name() string {
	return filepath.Base(i.Path)
}

// Ext liefert die Dateierweiterung des Bildes
func (i *Image) Ext() string {
	return filepath.Ext(i.Path)
}

// MimeType liefert entsprechend der Dateiendung den Mime Typ
func (i *Image) MimeType() string {
	return mime.TypeByExtension(i.Ext())
}

// Width gibt die Breite des Bildes per image.DecodeConfig zurück.
// Diese Funktion ist wesentlich performanter als die Breite über
// den image.Image Typ und desen Bounds() Methode zu holen
func (i *Image) Width() int {
	config, _, _ := image.DecodeConfig(i.Open())
	return config.Width
}

// Open öffnet die Datei und gibt einen Datei Handler zurück
func (i *Image) Open() *os.File {
	handler, err := os.Open(i.Path)
	if err != nil {
		panic(err)
	}
	return handler
}

// Image gibt eine image.Image Repräsentation des Bildes zurück
func (i *Image) Image() image.Image {
	image, err := imaging.Open(i.Path)
	if err != nil {
		panic(err)
	}
	return image
}

// Resize skaliert das Bild auf das angegebene Bildmaß
func (i Image) Resize(size ImageSize) {
	resized := imaging.Resize(i.Image(), size.Width, size.Height, imaging.Lanczos)
	dx, dy := resized.Bounds().Dx(), resized.Bounds().Dy()

	newFile := filepath.Join(i.Dir(), strconv.Itoa(dx)+"x"+strconv.Itoa(dy)+i.Ext())
	imaging.Save(resized, newFile)
}

// CropResize schneidet das Bild quadratisch zu und
// skaliert das Bild auf das angegebene Bildmaß
func (i Image) CropResize(size ImageSize) {
	var cropped *image.NRGBA
	image := i.Image()
	dx, dy := image.Bounds().Dx(), image.Bounds().Dy()

	if dx > dy {
		cropped = imaging.CropCenter(image, dy, dy)
	} else {
		cropped = imaging.CropCenter(image, dx, dx)
	}
	resized := imaging.Resize(cropped, size.Width, 0, imaging.Lanczos)

	dx, dy = resized.Bounds().Dx(), resized.Bounds().Dy()

	newFile := filepath.Join(i.Dir(), strconv.Itoa(dx)+"x"+strconv.Itoa(dy)+i.Ext())
	imaging.Save(resized, newFile)
}

// Quantize erstellt eine Farbpalette
func (i Image) Quantize(count int) ColorPalette {
	q := quantize.MedianCutQuantizer{}
	p := q.Quantize(make([]color.Color, 0, count), i.Image())
	fp := NewColorPalette(p)
	return fp
}

// SaveColorPalette speichert die quantisierte Farbpalette
func (i Image) SaveColorPalette(fileName string, palette ColorPalette) {
	byte, err := json.Marshal(palette)

	if err != nil {
		log.Panicf("Json Marshal fehlgeschlagen: %s", err.Error())
	}

	jsonFile, _ := os.Create(filepath.Join(i.Dir(), fileName))
	jsonFile.Write(byte)
	jsonFile.Close()
}
