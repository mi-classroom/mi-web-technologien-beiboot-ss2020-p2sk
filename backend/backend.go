package main

import (
	"errors"
	"html/template"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"./config"
	"./gallery"
	"github.com/gin-gonic/gin"
)

func toCSS(c color.NRGBA) template.CSS {
	s := "background-color: rgba(" + strings.Join([]string{strconv.Itoa(int(c.R)), strconv.Itoa(int(c.G)), strconv.Itoa(int(c.B)), strconv.Itoa(int(c.A))}, ", ") + ");"
	return template.CSS(s)
}

var (
	server *gin.Engine = gin.Default()
)

// Konfigurieren des Servers
func prepareServer() {
	// Bildverzeichnis
	server.Static("/uploads", config.UploadDir)

	server.SetFuncMap(template.FuncMap{
		"toCSS": toCSS,
	})

	// Templates
	server.LoadHTMLGlob(config.TemplateDir)
}

// Server
func main() {
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
				"uploaded": filepath.ToSlash(c.MustGet("image").(gallery.Picture).Path),
				"colors":   c.MustGet("colors"),
				"errors":   c.Errors,
			})
		},
	)

	// REST API
	v1 := server.Group("/rest/v1")
	{
		v1.Use(func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Next()
		})

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, time.Now())
		})

		/**
		Liefert alle CollectionIDs (Dirnames)
		@TODO JSON
		?count=int
		?sort=[alpha|date|color|random]
		*/
		v1.GET("/collections", func(c *gin.Context) {
			count, _ := strconv.Atoi(c.DefaultQuery("count", "10"))
			sort := c.DefaultQuery("sort", "alpha")

			galleryObj := gallery.LoadGallery(config.UploadDir, config.ColorFile, config.IgnoreFiles)
			galleryObj.Sort(gallery.SortType(sort))

			c.JSON(http.StatusOK, galleryObj.Reduce(count))
		})

		// Liefert alle Bilder einer Collection
		// @TODO JSON
		//v1.GET("/collections/:dirId")

		// Liefert ein bestimmtes Bild
		//v1.GET("/collections/:dirId/:pictureName")
	}

	server.Run()
}

// Endpunkt Bildübersicht
func overviewAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		images := gallery.LoadGallery(config.UploadDir, config.ColorFile, config.IgnoreFiles)
		c.HTML(http.StatusOK, "overview.tmpl", gin.H{
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

		path := gallery.CreatePictureFolder(config.UploadDir)

		file, _ := fileHeader.Open()
		defer file.Close()
		imageInfo, format, _ := image.DecodeConfig(file)

		// see https://github.com/gin-gonic/gin/issues/1693
		fileName := gallery.CreateFileName(imageInfo.Width, imageInfo.Height, format)
		fileDest := filepath.Join(path, fileName)

		if err := c.SaveUploadedFile(fileHeader, fileDest); err != nil {
			c.Error(err)
		}

		c.Set("image", gallery.Picture{Path: fileDest})

		c.Next()
		// Wenn im weiteren Verlauf ein Fehler auftritt, sollte
		// Datei ggf. vom Server gelöscht werden
	}
}

//
func scaleImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		image := c.MustGet("image").(gallery.Picture)

		defaultImageSizes := config.DefaultImageSizes

		// custom scaling
		scaleValue, scaleExists := c.Get("customScale")
		if scaleExists {
			scaleValue, _ = strconv.Atoi(scaleValue.(string))
			defaultImageSizes = append(defaultImageSizes, gallery.FromFactor(scaleValue.(int), image.Width()))
		}

		image.ProcessImageSizes(defaultImageSizes)

		c.Next()
	}
}

//
func quantizeImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		image := c.MustGet("image").(gallery.Picture)

		palette := image.SaveColorPalette(config.ColorFile, config.ColorCount)

		c.Set("colors", palette)
		c.Next()
	}
}
