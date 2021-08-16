package rot13reader

import "io"

type Rot13Reader struct {
	r io.Reader
}

func (r Rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i := range b {
		if (b[i] >= 65 && b[i] <= 77) || (b[i] >= 97 && b[i] <= 109) {
			b[i] += 13
		} else if (b[i] >= 78 && b[i] <= 90) || (b[i] >= 110 && b[i] <= 122) {
			b[i] -= 13
		}
	}
	return
}
