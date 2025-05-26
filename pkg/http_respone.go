package pkg

var message map[int]string

const SuccessCode = 200
const (
	BadRequestCode   = 400
	UnauthorizedCode = 401
	ForbiddenCode    = 403
	NotFoundCode     = 404
	InternalErrCode  = 500
)

func init() {
	message = map[int]string{}
	message[SuccessCode] = "Success"
	message[BadRequestCode] = "参数错误"
	message[UnauthorizedCode] = "未授权"
	message[ForbiddenCode] = "禁止访问"
	message[NotFoundCode] = "资源未找到"
	message[InternalErrCode] = "内部错误"
}
func Suceess(data any) map[string]any {
	m := map[string]any{}
	m["code"] = 200
	m["message"] = "Success"
	m["data"] = data
	return m
}
func Fail(data any) map[string]any {
	m := map[string]any{}
	m["code"] = 400
	m["message"] = "fail"
	m["data"] = data
	return m
}
func Error(data any) map[string]any {
	m := map[string]any{}
	m["code"] = 500
	m["message"] = "Error"
	m["data"] = data
	return m
}
