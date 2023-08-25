package xerr

import "fmt"

type CodeError struct {
	Code uint32
	Msg  string
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("CodeError: %d, %s", e.Code, e.Msg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{Code: errCode, Msg: errMsg}
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{Code: errCode, Msg: ErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{Code: ServerCommonError, Msg: errMsg}
}
