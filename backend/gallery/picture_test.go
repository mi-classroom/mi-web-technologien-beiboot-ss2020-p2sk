package gallery

import (
	"fmt"
	"testing"
)

const testImage = "logos.jpg"

var (
	imageDummy Picture = Picture{"testDir/testfile.png"}
)

func BenchmarkWidthPerOpenFile(b *testing.B) {
	picture := Picture{testImage}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		picture.Width()
	}
}

func BenchmarkWidthPerImage(b *testing.B) {
	picture := Picture{testImage}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp := picture.Image()
		tmp.Bounds().Dx()
	}
}

func ExamplePicture() {
	fmt.Println(imageDummy)

	picture := new(Picture)
	fmt.Print(picture)
	// Output: {testDir/testfile.png}
	// &{}
}

func TestDir(t *testing.T) {
	got := imageDummy.Dir()
	if got != "testDir" {
		t.Errorf("got Picture.Dir() = %s; want \"testDir\"", got)
	}
}

func TestName(t *testing.T) {
	got := imageDummy.Name()
	if got != "testfile.png" {
		t.Errorf("got Picture.Name() = %s; want \"testfile.png\"", got)
	}
}

func TestExt(t *testing.T) {
	got := imageDummy.Ext()
	if got != ".png" {
		t.Errorf("got Picture.Ext() = %s; want \".png\"", got)
	}
}

func TestMimeType(t *testing.T) {
	got := imageDummy.MimeType()
	if got != "image/png" {
		t.Errorf("got Picture.MimeType() = %s; want \"image/png\"", got)
	}
}

/*func TestJson(t *testing.T) {
	image := Image{testImage}
	q := image.Quantize(10)

	for _, c := range q {
		t.Log(c.(color.RGBA).R)
		t.Log(c.(color.RGBA).G)
		t.Log(c.(color.RGBA).B)
	}
	b, err := json.Marshal(q)

	if err != nil {
		t.Errorf("Bild.Name() = %s; want testfile.png", b)
	}

	q2 := make(color.Palette, 0, 0)
	//t.Errorf("Bild.Name() = %T; want testfile.png", q2)
	err = json.Unmarshal(b, &q2)
	t.Errorf("Bild.Name() = %v; want testfile.png", b)
	//log.Print(q2)
	// Output: 123
}*/

/*
func TestOpen(){}
func TestImage(){}
func TestResize(){}
*/
