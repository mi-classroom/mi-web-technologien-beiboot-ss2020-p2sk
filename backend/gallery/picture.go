package gallery

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strconv"

	"github.com/RobCherry/vibrant"
	"github.com/disintegration/imaging"
)

const fileTemplate = "%dx%d%s"

// CreateFileName erstellt einen Dateinamen
func CreateFileName(width, height int, format string) string {
	if format[:1] != "." {
		format = "." + format
	}
	return fmt.Sprintf(fileTemplate, width, height, format)
}

// Picture Model
type Picture struct {
	Path string `json:"path"`
}

func (p Picture) MarshalJSON() ([]byte, error) {
	jsonMap := map[string]string{
		"name":   p.Name(),
		"uri":    "todo",
		"width":  strconv.Itoa(p.Size().Width),
		"height": strconv.Itoa(p.Size().Height),
	}
	//print(p.Path)
	return json.Marshal(jsonMap)
}

// Dir liefert den Pfad des Bildes
func (p Picture) Dir() string {
	return filepath.Dir(p.Path)
}

// Name liefert den Name des Bildes
func (p Picture) Name() string {
	return filepath.Base(p.Path)
}

// Ext liefert die Dateierweiterung des Bildes
func (p Picture) Ext() string {
	return filepath.Ext(p.Path)
}

// MimeType liefert entsprechend der Dateiendung den Mime Typ
func (p Picture) MimeType() MimeType {
	return MimeType(mime.TypeByExtension(p.Ext()))
}

// Size liefert die Abmessung eines Bildes
func (p Picture) Size() ImageSize {
	config, _, _ := image.DecodeConfig(p.Open())
	return ImageSize{config.Width, config.Height}
}

// Width gibt die Breite des Bildes per image.DecodeConfig zurück.
// Diese Funktion ist wesentlich performanter als die Breite über
// den image.Image Typ und desen Bounds() Methode zu holen
func (p Picture) Width() int {
	config, _, _ := image.DecodeConfig(p.Open())
	return config.Width
}

// Open öffnet die Datei und gibt einen Datei Handler zurück
func (p Picture) Open() *os.File {
	handler, err := os.Open(p.Path)
	if err != nil {
		panic(err)
	}
	return handler
}

// Image gibt eine image.Image Repräsentation des Bildes zurück
func (p Picture) Image() image.Image {
	image, err := imaging.Open(p.Path)
	if err != nil {
		panic(err)
	}
	return image
}

// ProcessImageSizes
func (p Picture) ProcessImageSizes(imageSizes []ImageSize) {
	for _, imageSize := range imageSizes {
		if imageSize.IsQuad() {
			p.CropResize(imageSize)
		} else {
			p.Resize(imageSize)
		}
	}
}

// Resize skaliert das Bild auf das angegebene Bildmaß
func (p Picture) Resize(size ImageSize) {
	resized := imaging.Resize(p.Image(), size.Width, size.Height, imaging.Lanczos)
	newFile := filepath.Join(p.Dir(), CreateFileName(size.Width, size.Height, p.Ext()))
	imaging.Save(resized, newFile)
}

// CropResize schneidet das Bild quadratisch zu und
// skaliert das Bild auf das angegebene Bildmaß
func (p Picture) CropResize(size ImageSize) {
	var cropped *image.NRGBA
	currentImageSize := p.Size()
	//dx, dy := image.Bounds().Dx(), image.Bounds().Dy()

	if currentImageSize.Width > currentImageSize.Height {
		cropped = imaging.CropCenter(p.Image(), currentImageSize.Height, currentImageSize.Height)
	} else {
		cropped = imaging.CropCenter(p.Image(), currentImageSize.Width, currentImageSize.Width)
	}
	resized := imaging.Resize(cropped, size.Width, 0, imaging.Lanczos)

	//dx, dy = resized.Bounds().Dx(), resized.Bounds().Dy()

	newFile := filepath.Join(p.Dir(), CreateFileName(size.Width, size.Height, p.Ext()))
	imaging.Save(resized, newFile)
}

/*func saveImage(image image.Image, size ImageSize) {
	newFile := filepath.Join(p.Dir(), CreateFileName(size.Width, size.Height, p.Ext()))
	imaging.Save(resized, newFile)
}*/

// Quantize erstellt eine Farbpalette
func (p Picture) quantize(count int) ColorPalette {
	palette := vibrant.NewPaletteBuilder(p.Image()).
		MaximumColorCount(uint32(count)).
		Generate()

	var colorPalette ColorPalette

	for _, swatch := range palette.Swatches() {
		colorPalette = append(colorPalette, swatch.Color().(color.NRGBA))
	}

	/*sort.Sort(colorPalette, func() {

	})*/

	return colorPalette
}

// SaveColorPalette speichert die quantisierte Farbpalette
func (p Picture) SaveColorPalette(fileName string, colorCount int) ColorPalette {
	palette := p.quantize(colorCount)

	byte, err := json.Marshal(palette)

	if err != nil {
		log.Panicf("Json Marshal fehlgeschlagen: %s", err.Error())
	}

	jsonFile, _ := os.Create(filepath.Join(p.Dir(), fileName))
	jsonFile.Write(byte)
	jsonFile.Close()

	return palette
}
