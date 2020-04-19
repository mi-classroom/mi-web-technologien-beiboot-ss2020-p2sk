package main

import (
	"fmt"
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
	fmt.Print(bild)
	// Output: {testDir/testfile.png}
}

func ExampleNewBild() {
	bild := new(Bild)
	fmt.Print(bild)
	// Output: &{}
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

/*
func TestOpen(){}
func TestImage(){}
func TestResize(){}
*/
