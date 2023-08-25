package result

import "net/http"

type ResponseSuccessBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type ResponseErrorBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

type NullJSON struct{}

func Success(data any) *ResponseSuccessBean {
	return &ResponseSuccessBean{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: data,
	}
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{
		Code: errCode,
		Msg:  errMsg,
	}
}
