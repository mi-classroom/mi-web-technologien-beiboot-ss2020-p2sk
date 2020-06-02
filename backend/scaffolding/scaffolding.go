package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
	_ "os"
	"path/filepath"
	"strconv"

	"../config"
	"../gallery"
)

const picsumURI = "https://picsum.photos/v2/list?page=2&limit=20"

type picsumData []map[string]interface{}

func main() {
	body, _ := getURLData(picsumURI)
	var dat picsumData

	json.Unmarshal(body, &dat)

	for _, image := range dat {
		fmt.Println("Start downloading file: ", image["download_url"].(string))
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
	//print(format)
	//print(image)

	//err := ioutil.WriteFile("/tmp/tmp../uploads/test.jpg", body, 0644)

	if err != nil {
		panic(err)
	}

	return image
}

func saveImage(image image.Image) gallery.Image {
	width := strconv.Itoa(image.Bounds().Max.X)
	newDir := gallery.MakeImageDir("../" + config.UploadDir)
	fileName := "original-" + width + ".jpg"
	path := filepath.Join(newDir, fileName)
	file, _ := os.Create(path)

	jpeg.Encode(file, image, &jpeg.Options{Quality: 90})
	return gallery.Image{Path: path}
}
