package main

import (
	"fmt"
	"testing"
)

var (
	bild Bild = Bild{"testDir/testfile.png", BildTyp{"", 0}}
)

func ExampleBild() {
	fmt.Print(bild)
	// Output: {testDir/testfile.png { 0}}
}

func ExampleNewBild() {
	bild := new(Bild)
	fmt.Print(bild)
	// Output: &{ { 0}}
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
