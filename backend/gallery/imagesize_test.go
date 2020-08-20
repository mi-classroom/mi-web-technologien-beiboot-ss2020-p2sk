package gallery

import (
	"fmt"
	"testing"
)

func ExampleImageSize() {
	fmt.Println(ImageSize{})

	imageSize := new(ImageSize)
	fmt.Println(imageSize)

	fmt.Println(ImageSize{250, 372})
	// Output: {0 0}
	// &{0 0}
	// {250 372}
}

func TestIsQuad(t *testing.T) {
	imageSize := ImageSize{300, 300}
	got := imageSize.IsQuad()
	if !got {
		t.Errorf("got ImageSize.IsQuad() = %t ; want true", got)
	}
}

func TestFromFactor(t *testing.T) {
	got := FromFactor(30, 100)
	if got.Width != 30 {
		t.Errorf("got ImageSize.FromFactor().Width = %d; want 30", got)
	}
	if got.Height != 0 {
		t.Errorf("got ImageSize.FromFactor().Height = %d; want 0", got)
	}

	got = FromFactor(30, 275)
	if got.Width != 82 {
		t.Errorf("got ImageSize.FromFactor().Width = %d; want 82", got)
	}
}
