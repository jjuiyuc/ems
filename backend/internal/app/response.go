package app

import (
	"github.com/gin-gonic/gin"

	"der-ems/internal/e"
)

// Gin godoc
type Gin struct {
	*gin.Context
}

// Response godoc
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response returns standard response format
func (g *Gin) Response(httpCode, code int, data interface{}) {
	g.JSON(httpCode, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}
