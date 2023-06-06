package web

// ErrorCodeResp 错误码
type ErrorCodeResp struct {
	Code    string
	Message string
}

func (ec ErrorCodeResp) SetMsg(message string) {
	ec.Message = message
}
func (ec ErrorCodeResp) SetCode(code string) {
	ec.Code = code
}

var (
	SuccessCode = ErrorCodeResp{Code: "0", Message: "success"}
	FailedCode  = ErrorCodeResp{Code: "-1", Message: "failed"}
)
