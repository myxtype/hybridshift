package request

import (
	"frame/pkg/ecode"
	"net/http"
)

// 响应的数据结构
type ResponseJSON struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func (a *AppRequest) Response(err error, args ...interface{}) {
	var data interface{}
	if len(args) > 0 {
		data = args[0]
	}

	ec := ecode.Cause(err)
	a.c.JSON(http.StatusOK, &ResponseJSON{
		Code:    ec.Code(),
		Message: ec.Message(),
		Data:    data,
	})
}

func (a *AppRequest) AbortResponse(err error, args ...interface{}) {
	a.c.Abort()
	a.Response(err, args...)
}
