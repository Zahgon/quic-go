package http3

import (
	"bytes"
	"io"
	"net/http"
	"sync"

	"github.com/quic-go/qpack"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3/qlog"
	"github.com/quic-go/quic-go/qlogwriter"
)

const bodyCopyBufferSize = 8 * 1024

type requestWriter struct {
	mutex     sync.Mutex
	encoder   *qpack.Encoder
	headerBuf *bytes.Buffer
}

func newRequestWriter() *requestWriter { _ = "STUB: not implemented"; return nil }

func (w *requestWriter) WriteRequestHeader(wr io.Writer, req *http.Request, gzip bool, streamID quic.StreamID, qlogger qlogwriter.Recorder) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *requestWriter) writeHeaders(wr io.Writer, req *http.Request, gzip bool, streamID quic.StreamID, qlogger qlogwriter.Recorder) error {
	_ = "STUB: not implemented"
	return nil
}

func isExtendedConnectRequest(req *http.Request) bool { _ = "STUB: not implemented"; return false }

func (w *requestWriter) encodeHeaders(req *http.Request, addGzipHeader bool, trailers string, contentLength int64, doQlog bool) ([]qlog.HeaderField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func authorityAddr(authority string) (addr string) { _ = "STUB: not implemented"; return "" }

func validPseudoPath(v string) bool { _ = "STUB: not implemented"; return false }

func actualContentLength(req *http.Request) int64 { _ = "STUB: not implemented"; return 0 }

func shouldSendReqContentLength(method string, contentLength int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *requestWriter) WriteRequestTrailer(wr io.Writer, req *http.Request, streamID quic.StreamID, qlogger qlogwriter.Recorder) error {
	_ = "STUB: not implemented"
	return nil
}
