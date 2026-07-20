package http3

import (
	"compress/gzip"
	"io"
)

type gzipReader struct {
	body io.ReadCloser
	zr   *gzip.Reader
	zerr error
}

func newGzipReader(body io.ReadCloser) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (gz *gzipReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (gz *gzipReader) Close() error { _ = "STUB: not implemented"; return nil }
