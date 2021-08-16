package readers

import (
	"testing"

	"golang.org/x/tour/reader"
)

func TestReaders(t *testing.T) {
	reader.Validate(&MyReader{})
}
