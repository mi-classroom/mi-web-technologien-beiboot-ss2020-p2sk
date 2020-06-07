package gallery

import (
	"testing"
)

var (
	mimeType MimeType = MimeType("image/png")
)

func TestIsValid(t *testing.T) {
	if !IsValid(mimeType) {
		t.Errorf("got MimeType = %s; want \"image/png\"", mimeType)
	}
}

func TestUnValid(t *testing.T) {
	if IsValid("image/gif") {
		t.Errorf("MimeType = %s is not a valid type", "image/gif")
	}
}
