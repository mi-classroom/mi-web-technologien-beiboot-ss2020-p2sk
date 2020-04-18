package main

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

// Container
type Container struct {
	Dir    string
	Bilder []Bild
}

// ToScrset
func (c Container) ToScrset() string {
	var scrset []string
	for _, e := range c.Bilder {
		scrset = append(scrset, e.Pfad+" "+strconv.Itoa(e.Width())+"w")
	}
	return strings.Join(scrset, ", ")
}

const (
	uploadDir   = "uploads/"
	cssDir      = "css/"
	jsDir       = "js/"
	templateDir = "templates/**/*"
)

var (
	server *gin.Engine = gin.Default()
)

// Konfigurieren des Servers
func prepareServer() {
	// Statische Dateien wie JS, CSS
	server.Static("/css", cssDir)
	server.Static("/js", jsDir)
	// Bildverzeichnis
	server.Static("/uploads", uploadDir)
	//fs := http.Dir(uploadDir)
	//server.StaticFS("/uploads", fs)

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

	//server.Get("/bild/:dir",)

	// Bilder upload
	server.POST(
		"/upload",
		validiereUpload(),
		persistiereBild(),
		skaliereBild(),
		func(c *gin.Context) {
			c.HTML(http.StatusOK, "uploaded.tmpl", gin.H{
				"uploaded": filepath.ToSlash(c.MustGet("bild").(Bild).Pfad),
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

		if !isValid(file.Header.Get("Content-Type")) {
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
		dateiName := strings.Join([]string{"original-", strconv.Itoa(imageInfo.Width), filepath.Ext(fileHeader.Filename)}, "")

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

func readImagesFromDir(dir string) map[string]Container {
	var lastDir string
	ignoreFiles := []string{".gitkeep"}
	container := make(map[string]Container)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == filepath.Dir(uploadDir) {
			return err
		}

		if info.IsDir() {
			lastDir = info.Name()
			container[lastDir] = Container{lastDir, make([]Bild, 0)}
			return nil
		}

		if stringInSlice(ignoreFiles, filepath.Base(path)) {
			return nil
		}
		// Pfadseperator / Unix/Windows
		path = filepath.ToSlash(path)

		tmpContainer := container[lastDir]
		tmpContainer.Bilder = append(tmpContainer.Bilder, Bild{path})
		container[lastDir] = tmpContainer
		return nil
	})

	return container
}
