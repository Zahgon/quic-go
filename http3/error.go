package http3

type Error struct {
	Remote       bool
	ErrorCode    ErrCode
	ErrorMessage string
}

var _ error = &Error{}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Is(target error) bool { _ = "STUB: not implemented"; return false }

func maybeReplaceError(err error) error { _ = "STUB: not implemented"; return nil }
