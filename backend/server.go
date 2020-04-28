package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

func stringInSlice(s []string, needle string) bool {
	for _, item := range s {
		if item == needle {
			return true
		}
	}
	return false
}

func toCSS(c color.RGBA) template.CSS {
	s := "background-color: rgba(" + strings.Join([]string{strconv.Itoa(int(c.R)), strconv.Itoa(int(c.G)), strconv.Itoa(int(c.B)), strconv.Itoa(int(c.A))}, ", ") + ");"
	return template.CSS(s)
}

var (
	server *gin.Engine = gin.Default()
)

// Konfigurieren des Servers
func prepareServer() {
	// Bildverzeichnis
	server.Static("/uploads", uploadDir)

	server.SetFuncMap(template.FuncMap{
		"toCSS": toCSS,
	})

	// Templates
	server.LoadHTMLGlob(templateDir)
}

// Server
func main() {
	fmt.Println("Golang Backendkomponente MI-Beibootprojekt")
	fmt.Println("Picturebox")

	prepareServer()

	// Übersichtseite ausliefern
	server.GET(
		"/",
		liefereBilderAction(),
	)

	// Bilder upload
	server.POST(
		"/upload",
		validiereUpload(),
		persistiereBild(),
		skaliereBild(),
		quantisiereBild(),
		func(c *gin.Context) {
			c.HTML(http.StatusOK, "uploaded.tmpl", gin.H{
				"uploaded": filepath.ToSlash(c.MustGet("bild").(Bild).Pfad),
				"farben":   c.MustGet("farben"),
				"errors":   c.Errors,
			})
		},
	)

	server.Run()
}

// Endpunkt Bildübersicht
func liefereBilderAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		bilder := readImagesFromDir(uploadDir)
		c.HTML(http.StatusOK, "overview.tmpl", gin.H{
			"title":  "Übersichtsseite",
			"bilder": bilder,
		})
	}
}

// Middleware zur Prüfung auf valide Datei
func validiereUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("skalierung", c.PostForm("skalierung"))

		file, err := c.FormFile("image")

		if err != nil {
			c.Error(errors.New("Ihr Bild hat keinen validen Typ"))
		}

		if !isValid(MimeType(file.Header.Get("Content-Type"))) {
			c.Error(errors.New("Invalider Mediatyp"))
			c.AbortWithStatus(http.StatusUnsupportedMediaType)
			return
		}

		c.Set("image", file)

		c.Next()
	}
}

//
func persistiereBild() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader := c.MustGet("image").(*multipart.FileHeader)
		ordnerName := time.Now().Format("20060102150405")

		err := os.Mkdir(filepath.Join(uploadDir, ordnerName), 0755)

		if err != nil {
			c.Error(err)
		}

		file, _ := fileHeader.Open()
		defer file.Close()
		imageInfo, _, _ := image.DecodeConfig(file)

		// see https://github.com/gin-gonic/gin/issues/1693
		dateiName := strings.Join([]string{strconv.Itoa(imageInfo.Width), "x", strconv.Itoa(imageInfo.Height), strings.ToLower(filepath.Ext(fileHeader.Filename))}, "")

		dateiDest := filepath.Join(uploadDir, ordnerName, dateiName)
		if err := c.SaveUploadedFile(fileHeader, dateiDest); err != nil {
			c.Error(err)
		}

		c.Set("bild", Bild{dateiDest})

		c.Next()
		// Wenn im weiteren Verlauf ein Fehler auftritt, sollte
		// Datei ggf. vom Server gelöscht werden
	}
}

//
func skaliereBild() gin.HandlerFunc {
	return func(c *gin.Context) {
		bild := c.MustGet("bild").(Bild)
		sFaktor, sExists := c.Get("skalierung")
		sFaktor, _ = strconv.Atoi(sFaktor.(string))

		for _, v := range defaultMaße {
			if v.isQuad() {
				bild.CropResize(v)
			} else {
				bild.Resize(v)
			}
		}

		// custom
		if sExists {
			customMaß := perFaktor(sFaktor.(int), bild.Image().Bounds().Dx())
			bild.Resize(customMaß)
		}

		c.Next()
	}
}

//
func quantisiereBild() gin.HandlerFunc {
	return func(c *gin.Context) {
		bild := c.MustGet("bild").(Bild)
		palette := bild.Quantisiere(anzahlFarben)

		bild.SpeicherColorMap(palette, colorFile)

		c.Set("farben", palette)
		c.Next()
	}
}

// Liest alle Bilder aus dem Upload Verzeichnis
func readImagesFromDir(dir string) map[string]Galerie {
	var lastDir string
	galerie := make(map[string]Galerie)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == filepath.Dir(uploadDir) {
			return err
		}

		if info.IsDir() {
			lastDir = info.Name()
			galerie[lastDir] = Galerie{lastDir, Farbpalette{}, make([]Bild, 0)}
			return nil
		}

		if stringInSlice(ignoreFiles, filepath.Base(path)) {
			return nil
		}

		if info.Name() == colorFile {
			var rgba Farbpalette
			tmp := galerie[lastDir]
			data, _ := ioutil.ReadFile(path)

			json.Unmarshal(data, &rgba)
			tmp.ColorMap = rgba
			galerie[lastDir] = tmp
		}
		// Pfadseperator / Unix/Windows
		path = filepath.ToSlash(path)

		tmpGalerie := galerie[lastDir]
		tmpGalerie.Sammlung = append(tmpGalerie.Sammlung, Bild{path})
		galerie[lastDir] = tmpGalerie
		return nil
	})

	return galerie
}
