package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

type HttpErr struct {
	ErrNo    ErrNo
	HttpCode int64
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func NewTokenErr(code int32, msg string) error {
	return &ErrNo{ErrCode: code, ErrMsg: msg}
}

func NewHttpErr(code int, httpCode int, msg string) HttpErr {
	return HttpErr{
		ErrNo:    ErrNo{ErrCode: int32(code), ErrMsg: msg},
		HttpCode: int64(httpCode),
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
