package web

import (
	"time"
)

// ApiResponse 统一接口响应结构
type ApiResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	NowTs   int64  `json:"now_ts"`
}

func NewBaseResponse(errCode ErrorCodeResp) *ApiResponse {
	return &ApiResponse{
		Code:    errCode.Code,
		Message: errCode.Message,
		Data:    make(map[string]any, 0),
		NowTs:   time.Now().Unix(),
	}
}

// SetData 设置响应数据
func (br *ApiResponse) SetData(respData any) *ApiResponse {
	br.Data = respData
	return br
}

// SuccessResp 成功的响应
func SuccessResp(respData any) *ApiResponse {
	if respData == nil {
		respData = make(map[string]any, 0)
	}
	return NewBaseResponse(SuccessCode).SetData(respData)
}

// FailResp 失败的响应
// code 默认为 FailedCode
// message 可自定义
func FailResp(message string, errCode ...ErrorCodeResp) *ApiResponse {
	_errCode := FailedCode
	if len(errCode) > 0 {
		_errCode = errCode[0]
	}
	_errCode.Message = message
	return NewBaseResponse(_errCode)
}

// FailRespWithError 失败的响应
// code 默认为 FailedCode
func FailRespWithError(errCode ...ErrorCodeResp) *ApiResponse {
	_errCode := FailedCode
	if len(errCode) > 0 {
		_errCode = errCode[0]
	}
	return NewBaseResponse(_errCode)
}
