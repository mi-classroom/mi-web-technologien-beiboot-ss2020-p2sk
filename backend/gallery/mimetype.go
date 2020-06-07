package gallery

// MimeType beschreibt das Konzept des MimeType
type MimeType string

// ValidMimes fasst die gültigen Image Mimes zusammen
var ValidMimes []MimeType = []MimeType{
	"image/jpeg",
	"image/png",
}

// IsValid prüft ob ein gegebenes MimeType gültig ist
func IsValid(mimeType MimeType) bool {
	for _, mime := range ValidMimes {
		if mime == mimeType {
			return true
		}
	}
	return false
}
