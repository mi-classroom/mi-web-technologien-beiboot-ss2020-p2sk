package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
	_ "os"
	"path"
	"path/filepath"

	"../config"
	"../gallery"
)

const uploadDir = "../" + config.UploadDir
const defaultQuality = 90

var picsumURI = "https://picsum.photos/v2/list?page=%d&limit=%d"

var page = flag.Int("p", 1, "the page number to download from")
var count = flag.Int("c", 20, "the image count to scaffold")
var force = flag.Bool("f", false, "force the download if upload dir has images")
var delete = flag.Bool("d", false, "DELETE all images from the upload dir")

type picsumData []map[string]interface{}

func init() {
	flag.Parse()
	picsumURI = fmt.Sprintf(picsumURI, *page, *count)
}

func main() {
	if *delete {
		deleteImages()
	}

	if hasImages() && !*force {
		fmt.Print("There are already images. Aborting.")
		os.Exit(1)
	}

	fmt.Printf("Getting %d images. \n", *count)

	body, _ := getURLData(picsumURI)
	var dat picsumData

	json.Unmarshal(body, &dat)

	for i, image := range dat {
		downloadURL := image["download_url"].(string)
		fmt.Printf("Start downloading file %d: %s \n", i+1, downloadURL)

		savedImage := saveImage(downloadURL)
		savedImage.ProcessImageSizes(config.DefaultImageSizes)

		fmt.Println("Quantize image")
		savedImage.SaveColorPalette(config.ColorFile, config.ColorCount)
	}
}

func deleteImages() {
	fmt.Println("Should I delete all images from the upload dir? [y/n]")

	answer := readCharFromStdin()

	switch answer {
	case 'y', 'Y':
		dirItems, _ := ioutil.ReadDir(uploadDir)
		for _, dirItem := range dirItems {
			if dirItem.IsDir() {
				os.RemoveAll(path.Join([]string{uploadDir, dirItem.Name()}...))
			}
		}
		os.Exit(0)
	default:
		fmt.Println("Nothing deleted.")
		os.Exit(0)
	}
}

func readCharFromStdin() rune {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return char
}

func hasImages() bool {
	dirItems, _ := ioutil.ReadDir(uploadDir)

	for _, dirItem := range dirItems {
		if dirItem.IsDir() {
			return true
		}
	}
	return false
}

func getURLData(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func saveImage(url string) gallery.Picture {
	body, err := getURLData(url)

	if err != nil {
		panic(err)
	}

	imageConfig, format, _ := image.DecodeConfig(bytes.NewReader(body))

	fileName := gallery.CreateFileName(imageConfig.Width, imageConfig.Height, format)
	imageDir := gallery.CreatePictureFolder(uploadDir)
	path := filepath.Join(imageDir, fileName)

	image, _, err := image.Decode(bytes.NewReader(body))

	if err != nil {
		panic(err)
	}

	file, _ := os.Create(path)
	jpeg.Encode(file, image, &jpeg.Options{Quality: defaultQuality})
	return gallery.Picture{Path: path}
}
