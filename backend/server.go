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

	"github.com/disintegration/imaging"

	"github.com/gin-gonic/gin"
)

type validMime struct {
	mimes []string
}

func (mime *validMime) isValid(mimeType string) bool {
	for _, mime := range mime.mimes {
		if mime == mimeType {
			return true
		}
	}
	return false
}

type BildTyp struct {
	Typ  string
	Size int
}

func FromFile(typen []BildTyp, pfad string) BildTyp {
	dateiName := filepath.Base(pfad)

	for _, item := range typen {
		if strings.Contains(dateiName, item.Typ) {
			return item
		}
	}
	tmp := strings.Split(strings.Split(dateiName, ".")[0], "-")
	typ := tmp[0]
	size, _ := strconv.Atoi(tmp[1])
	return BildTyp{typ, size}
}

func stringInSlice(s []string, needle string) bool {
	for _, item := range s {
		if item == needle {
			return true
		}
	}
	return false
}

type Container struct {
	Dir    string
	Bilder []Bild
}

func (c Container) ToScrset() string {
	var scrset []string
	for _, e := range c.Bilder {
		scrset = append(scrset, e.Pfad+" "+strconv.Itoa(e.Typ.Size)+"w")
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
	validMimeTypes = validMime{[]string{"image/png", "image/jpeg"}}
	bildTypen      = []BildTyp{
		BildTyp{"desktop", 768},
		BildTyp{"tablet", 640},
		BildTyp{"quad", 400},
		BildTyp{"mobile", 360},
	}
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
	fmt.Println("Feature: Image Upload")

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
				"uploaded": filepath.ToSlash(c.MustGet("dateiPfad").(string)),
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
		tmpContainer.Bilder = append(tmpContainer.Bilder, Bild{path, FromFile(bildTypen, path)})
		container[lastDir] = tmpContainer
		return nil
	})

	return container
}

// Middleware zur Prüfung auf valide Datei
func validiereUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("skalierung", c.PostForm("skalierung"))

		file, err := c.FormFile("image")

		if err != nil {
			c.Error(errors.New("Ihr Bild hat keinen validen Typ"))
		}

		if !validMimeTypes.isValid(file.Header.Get("Content-Type")) {
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

		err := os.Mkdir(filepath.Join(uploadDir, ordnerName), os.ModeDir)

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
		c.Set("dateiPfad", dateiDest)

		c.Next()
		// Wenn im weiteren Verlauf ein Fehler auftritt, sollte
		// Datei ggf. vom Server gelöscht werden
	}
}

//
func skaliereBild() gin.HandlerFunc {
	return func(c *gin.Context) {
		dateiPfad := c.MustGet("dateiPfad").(string)
		sFaktor, sExists := c.Get("skalierung")
		sFaktor, _ = strconv.Atoi(sFaktor.(string))
		bild, err := imaging.Open(dateiPfad)

		if err != nil {
			c.Error(errors.New("Bild konnte nicht geöffnet werden"))
		}

		pfad := filepath.Dir(dateiPfad)
		ext := filepath.Ext(dateiPfad)

		for _, v := range bildTypen {
			var resized *image.NRGBA
			if v.Typ == "quad" {
				var cropped *image.NRGBA
				dx := bild.Bounds().Dx()
				dy := bild.Bounds().Dy()

				if dx > dy {
					cropped = imaging.CropCenter(bild, dy, dy)
				} else {
					cropped = imaging.CropCenter(bild, dx, dx)
				}

				resized = imaging.Resize(cropped, v.Size, 0, imaging.Lanczos)
			} else {
				resized = imaging.Resize(bild, v.Size, 0, imaging.Lanczos)
			}
			imaging.Save(resized, filepath.Join(pfad, v.Typ+"-"+strconv.Itoa(v.Size)+ext))
		}

		// custom
		if sExists {
			width := bild.Bounds().Dx() * sFaktor.(int) / 100
			print(width)
			resized := imaging.Resize(bild, width, 0, imaging.Lanczos)
			pfad = filepath.Join(pfad, "custom-"+strconv.Itoa(width)+ext)
			imaging.Save(resized, pfad)
		}

		c.Next()
	}
}
