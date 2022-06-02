package app

import (
	"github.com/gin-gonic/gin"

	"der-ems/internal/e"
)

type Gin struct {
	*gin.Context
}

// Standard response format
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, code int, data interface{}) {
	g.JSON(httpCode, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}
