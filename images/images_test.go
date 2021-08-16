package images

import (
	"testing"

	"golang.org/x/tour/pic"
)

func TestImages(t *testing.T) {
	m := Image{255, 255}
	pic.ShowImage(m)
}
