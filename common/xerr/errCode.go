package xerr

const (
	_ uint32 = iota
	commonPrefix
)

const (
	_ = commonPrefix*1000 + iota
	ServerCommonError
	RequestParamError
)
