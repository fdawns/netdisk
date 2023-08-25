package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[RequestParamError] = "请求参数错误"
}

func ErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	}
	return message[ServerCommonError]
}
