package http3

import (
	"errors"
	"io"
	"net/http"

	"github.com/quic-go/qpack"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
)

type qpackError struct{ err error }

func (e *qpackError) Error() string { _ = "STUB: not implemented"; return "" }
func (e *qpackError) Unwrap() error { _ = "STUB: not implemented"; return nil }

var errHeaderTooLarge = errors.New("http3: headers too large")

type header struct {
	Path      string
	Method    string
	Authority string
	Scheme    string
	Status    string

	Protocol string

	ContentLength int64

	Headers http.Header
}

var invalidHeaderFields = [...]string{
	"connection",
	"keep-alive",
	"proxy-connection",
	"transfer-encoding",
	"upgrade",
}

func parseHeaders(decodeFn qpack.DecodeFunc, isRequest bool, sizeLimit int, headerFields *[]qpack.HeaderField) (header, error) {
	_ = "STUB: not implemented"
	return *new(header), nil
}

func validateHeaderFieldNameAndValue(h qpack.HeaderField) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRegularHeaderField(h qpack.HeaderField) error { _ = "STUB: not implemented"; return nil }

func validateTrailerHeaderField(h qpack.HeaderField) error { _ = "STUB: not implemented"; return nil }

func parseTrailers(decodeFn qpack.DecodeFunc, sizeLimit int, headerFields *[]qpack.HeaderField) (http.Header, error) {
	_ = "STUB: not implemented"
	return *new(http.Header), nil
}

func requestFromHeaders(decodeFn qpack.DecodeFunc, sizeLimit int, headerFields *[]qpack.HeaderField) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validExtendedConnectProtocol(protocol string) bool { _ = "STUB: not implemented"; return false }

func updateResponseFromHeaders(rsp *http.Response, decodeFn qpack.DecodeFunc, sizeLimit int, headerFields *[]qpack.HeaderField) error {
	_ = "STUB: not implemented"
	return nil
}

func extractAnnouncedTrailers(header http.Header) http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func writeTrailers(wr io.Writer, trailers http.Header, streamID quic.StreamID, qlogger qlogwriter.Recorder) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func decodeTrailers(r io.Reader, hf *headersFrame, maxHeaderBytes int, decoder *qpack.Decoder, qlogger qlogwriter.Recorder, streamID quic.StreamID) (http.Header, error) {
	_ = "STUB: not implemented"
	return *new(http.Header), nil
}
