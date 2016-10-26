package error

import "fmt"

const (
	OK                    = 0

	ERROR_REQUEST_INVALID = 10
	ERROR_MISSING_PARAM = 11

	ERROR_ANTI_FRAUD_BUYER = 100

	ERROR_UNKNOWN = 31415
)

func ReturnMessage(code int, args ...interface{}) string {
	switch code {
	case ERROR_REQUEST_INVALID:
		return fmt.Sprintf("请求有误:%s", args...)
	case ERROR_MISSING_PARAM:
		return fmt.Sprintf("参数有误")
	case OK:
		return "OK"
	default:
		return "未知"
	}
}

