package serializer

import "github.com/gin-gonic/gin"

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误

const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

// Response 基础序列化器
type Response struct {
	Code     int    `json:"code"`
	Success  bool   `json:"success"`
	ErrorMsg string `json:"errorMsg"`
	Data     any    `json:"data"`
	Total    int    `json:"total"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Success:  false,
		Code:     errCode,
		ErrorMsg: msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.ErrorMsg = err.Error()
	}
	return res
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}
func Success(data any) Response {
	res := Response{
		Success: true,
		Data:    data,
	}
	return res
}

func CheckLogin() Response {
	return Err(CodeCheckLogin, "未登录", nil)
}
