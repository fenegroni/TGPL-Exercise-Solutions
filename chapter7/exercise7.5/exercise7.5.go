package exercise7_5

import "io"

type LimitedReader struct {
	r        io.Reader
	limit, n int64
}

func (lr *LimitedReader) Read(p []byte) (int, error) {
	if lr.n == lr.limit {
		return 0, io.EOF
	}
	q := p[:lr.limit-lr.n]
	n, err := lr.r.Read(q)
	lr.n += int64(n)
	if err == nil && lr.n == lr.limit {
		err = io.EOF
	}
	return n, err
}

func LimitReader(r io.Reader, limit int64) io.Reader {
	var lr = LimitedReader{r, limit, 0}
	return &lr
}
