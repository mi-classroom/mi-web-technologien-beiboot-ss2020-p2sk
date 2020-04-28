package main

import (
	"fmt"
	"image/color"
	"testing"
)

var (
	bild Bild = Bild{"testDir/testfile.png"}
)

func BenchmarkWidthPerOpenFile(b *testing.B) {
	bild := Bild{"logos.jpg"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bild.Width()
	}
}

func BenchmarkWidthPerImage(b *testing.B) {
	bild := Bild{"logos.jpg"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp := bild.Image()
		tmp.Bounds().Dx()
	}
}

func ExampleBild() {
	fmt.Println(bild)

	bild := new(Bild)
	fmt.Print(bild)
	// Output: {testDir/testfile.png}
	// &{}
}

func TestDir(t *testing.T) {
	got := bild.Dir()
	if got != "testDir" {
		t.Errorf("Bild.Dir() = %s; want testDir", got)
	}
}

func TestName(t *testing.T) {
	got := bild.Name()
	if got != "testfile.png" {
		t.Errorf("Bild.Name() = %s; want testfile.png", got)
	}
}

func TestExt(t *testing.T) {
	got := bild.Ext()
	if got != ".png" {
		t.Errorf("Bild.Ext() = %s; want .png", got)
	}
}

func TestMimeType(t *testing.T) {
	got := bild.MimeType()
	if got != "image/png" {
		t.Errorf("Bild.MimeType() = %s; want image/png", got)
	}
}

func TestJson(t *testing.T) {
	bild := Bild{"logos.jpg"}
	q := bild.Quantisiere()

	for _, c := range q {
		t.Log(c.(color.RGBA).R)
		t.Log(c.(color.RGBA).G)
		t.Log(c.(color.RGBA).B)
	}

	/*b, err := json.Marshal(q)

	if err != nil {
		t.Errorf("Bild.Name() = %s; want testfile.png", b)
	}

	q2 := make(color.Palette, 0, 0)
	//t.Errorf("Bild.Name() = %T; want testfile.png", q2)
	err = json.Unmarshal(b, &q2)
	t.Errorf("Bild.Name() = %v; want testfile.png", b)
	//log.Print(q2)
	// Output: 123*/
}

/*
func TestOpen(){}
func TestImage(){}
func TestResize(){}
*/
