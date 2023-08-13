package zerror

import "fmt"

type ZError struct {
	code uint32
	msg  string
}

func (e *ZError) Error() string {
	return fmt.Sprintf("{code:%d, errMsg:%s}", e.code, e.msg)
}

func (e *ZError) GetCode() uint32 {
	return e.code
}
func (e *ZError) GetMsg() string {
	return e.msg
}

func NewZError(code uint32, errMsg string) *ZError {
	return &ZError{
		code: code,
		msg:  errMsg,
	}
}
func NewZErrorCode(code uint32) *ZError {
	c := SERVER_COMMON_ERROR
	if _, ok := message[code]; ok {
		c = code
	}
	return &ZError{
		code: c,
		msg:  message[c],
	}
}
