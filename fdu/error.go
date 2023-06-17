package fdu

type ErrorCode int

const (
	ErrNetwork ErrorCode = iota
	ErrParseHTML
	ErrParseJSON
	ErrNeedCaptcha
	ErrInvalidPassword
	ErrLogin
)

func (e ErrorCode) String() string {
	switch e {
	case ErrNetwork:
		return "network error"
	case ErrParseHTML:
		return "parse html error"
	case ErrParseJSON:
		return "parse json error"
	case ErrNeedCaptcha:
		return "need captcha"
	case ErrInvalidPassword:
		return "invalid password"
	case ErrLogin:
		return "login failed"
	default:
		return "unknown error"
	}
}

type Error struct {
	Code ErrorCode
	Msg  string
}

func (e *Error) Error() string {
	return e.Code.String() + ": " + e.Msg
}

func FDUErr(code ErrorCode, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
