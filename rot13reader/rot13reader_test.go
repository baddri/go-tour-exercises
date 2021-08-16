package rot13reader

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestRot13reader(t *testing.T) {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
