package readers

type MyReader struct{}

func (r *MyReader) Read(b []byte) (n int, err error) {
	n = 0
	for ; n < len(b); n++ {
		b[n] = 65
	}
	return
}
