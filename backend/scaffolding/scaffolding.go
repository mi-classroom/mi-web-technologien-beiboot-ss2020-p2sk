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
	"strconv"

	"../config"
	"../gallery"
)

const uploadDir = "../" + config.UploadDir

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
		os.Exit(0)
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
		fmt.Printf("Start downloading file %d: %s \n", i+1, image["download_url"].(string))
		dImage := downloadImage(image["download_url"].(string))
		sImage := saveImage(dImage)

		for _, s := range config.DefaultImageSizes {
			if s.IsQuad() {
				sImage.CropResize(s)
			} else {
				sImage.Resize(s)
			}
		}
	}
}

func deleteImages() {
	fmt.Println("Should I delete all images from the upload dir? [y/n]")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	fmt.Println(char == 'y')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch char {
	case 'y', 'Y':
		dir, _ := ioutil.ReadDir(uploadDir)
		for _, d := range dir {
			if d.IsDir() {
				os.RemoveAll(path.Join([]string{uploadDir, d.Name()}...))
			}
		}
		os.Exit(0)
		break
	default:
		fmt.Println("Nothing deleted.")
		os.Exit(0)
	}

}

func hasImages() bool {
	dirItems, _ := ioutil.ReadDir(uploadDir)

	for _, item := range dirItems {
		if item.IsDir() {
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

func downloadImage(url string) image.Image {
	body, _ := getURLData(url)

	image, _, err := image.Decode(bytes.NewReader(body))

	if err != nil {
		panic(err)
	}

	return image
}

func saveImage(image image.Image) gallery.Image {
	width := strconv.Itoa(image.Bounds().Max.X)
	newDir := gallery.MakeImageDir(uploadDir)
	fileName := "original-" + width + ".jpg"
	path := filepath.Join(newDir, fileName)
	file, _ := os.Create(path)

	jpeg.Encode(file, image, &jpeg.Options{Quality: 90})
	return gallery.Image{Path: path}
}
