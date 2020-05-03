package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
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
	"time"

	"./gallery"
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
		overviewAction(),
	)

	// Bilder upload
	server.POST(
		"/upload",
		validateUpload(),
		persistImage(),
		scaleImage(),
		quantizeImage(),
		func(c *gin.Context) {
			c.HTML(http.StatusOK, "uploaded.tmpl", gin.H{
				"uploaded": filepath.ToSlash(c.MustGet("image").(gallery.Image).Path),
				"colors":   c.MustGet("colors"),
				"errors":   c.Errors,
			})
		},
	)

	server.Run()
}

// Endpunkt Bildübersicht
func overviewAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		images := readImagesFromDir(uploadDir)
		c.HTML(http.StatusOK, "overview.tmpl", gin.H{
			"title":  "Übersichtsseite",
			"images": images,
		})
	}
}

// Middleware zur Prüfung auf valide Datei
func validateUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("customScale", c.PostForm("customScale"))

		file, err := c.FormFile("image")

		if err != nil {
			c.Error(errors.New("Ihr Bild hat keinen validen Typ"))
		}

		if !gallery.IsValid(gallery.MimeType(file.Header.Get("Content-Type"))) {
			c.Error(errors.New("Invalider Mediatyp"))
			c.AbortWithStatus(http.StatusUnsupportedMediaType)
			return
		}

		c.Set("image", file)

		c.Next()
	}
}

//
func persistImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader := c.MustGet("image").(*multipart.FileHeader)
		dirName := time.Now().Format("20060102150405")

		err := os.Mkdir(filepath.Join(uploadDir, dirName), 0755)

		if err != nil {
			c.Error(err)
		}

		file, _ := fileHeader.Open()
		defer file.Close()
		imageInfo, _, _ := image.DecodeConfig(file)

		// see https://github.com/gin-gonic/gin/issues/1693
		fileName := strings.Join([]string{strconv.Itoa(imageInfo.Width), "x", strconv.Itoa(imageInfo.Height), strings.ToLower(filepath.Ext(fileHeader.Filename))}, "")

		fileDest := filepath.Join(uploadDir, dirName, fileName)
		if err := c.SaveUploadedFile(fileHeader, fileDest); err != nil {
			c.Error(err)
		}

		c.Set("image", gallery.Image{fileDest})

		c.Next()
		// Wenn im weiteren Verlauf ein Fehler auftritt, sollte
		// Datei ggf. vom Server gelöscht werden
	}
}

//
func scaleImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		image := c.MustGet("image").(gallery.Image)

		for _, s := range defaultImageSizes {
			if s.IsQuad() {
				image.CropResize(s)
			} else {
				image.Resize(s)
			}
		}

		// custom scaling
		scaleValue, scaleExists := c.Get("customScale")
		if scaleExists {
			scaleValue, _ = strconv.Atoi(scaleValue.(string))
			customSize := gallery.FromFactor(scaleValue.(int), image.Width() /*Image().Bounds().Dx()*/)
			image.Resize(customSize)
		}

		c.Next()
	}
}

//
func quantizeImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		image := c.MustGet("image").(gallery.Image)
		palette := image.Quantize(colorCount)

		image.SaveColorPalette(colorFile, palette)

		c.Set("colors", palette)
		c.Next()
	}
}

// Liest alle Bilder aus dem Upload Verzeichnis
func readImagesFromDir(dir string) map[string]gallery.Gallery {
	var lastDir string
	tmpGallery := make(map[string]gallery.Gallery)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == filepath.Dir(uploadDir) {
			return err
		}

		if info.IsDir() {
			lastDir = info.Name()
			tmpGallery[lastDir] = gallery.Gallery{lastDir, gallery.ColorPalette{}, make([]gallery.Image, 0)}
			return nil
		}

		if stringInSlice(ignoreFiles, filepath.Base(path)) {
			return nil
		}

		if info.Name() == colorFile {
			var rgba gallery.ColorPalette
			tmp := tmpGallery[lastDir]
			data, _ := ioutil.ReadFile(path)

			json.Unmarshal(data, &rgba)
			tmp.ColorMap = rgba
			tmpGallery[lastDir] = tmp
		}
		// Pfadseperator / Unix/Windows
		path = filepath.ToSlash(path)

		tmpGalerie := tmpGallery[lastDir]
		tmpGalerie.Collection = append(tmpGalerie.Collection, gallery.Image{path})
		tmpGallery[lastDir] = tmpGalerie
		return nil
	})

	return tmpGallery
}
