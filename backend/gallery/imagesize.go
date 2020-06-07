package gallery

// ImageSize repräsentiert ein Bildmaß aus Breite und Höhe
type ImageSize struct {
	Width  int
	Height int
}

// IsQuad
func (is ImageSize) IsQuad() bool {
	return is.Width == is.Height
}

// FromFactor
func FromFactor(factor, width int) ImageSize {
	return ImageSize{factor * width / 100, 0}
}
