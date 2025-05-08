package errno

import (
	"errors"
	"fmt"
)

type Errno struct {
	Code    int64
	Message string
}

func NewErrno(code int64, meg string) Errno {
	return Errno{
		Code:    code,
		Message: meg,
	}
}

func (e Errno) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e Errno) WithMessage(msg string) Errno {
	e.Message = msg
	return e
}

func (e Errno) WithFormat(format string, a ...any) Errno {
	e.Message = fmt.Sprintf(format, a)
	return e
}

func ConvertErr(err error) Errno {
	errno := Errno{}
	if err == nil {
		return Success
	}
	if errors.As(err, &errno) {
		return errno
	}
	s := SystemErr
	s.Message = err.Error()
	return s
}
