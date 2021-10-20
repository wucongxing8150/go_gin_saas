package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	Ctx *gin.Context
}

/**
 * @title Response
 * @Description: 返回值结构体
 **/
type Response struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

/**
 * @title Return
 * @Description: 自定义返回 可兼容失败成功 用于其他code值返回
 * @param code code值
 * @param message 返回信息
 * @param data 返回数据值
 * @return *Response 上方定义的Response结构体
 * @Date 2021-10-20 14:31:54
 **/
func BusinessReturn(code int, message string, data interface{}) *Response {
	return &Response{
		Code : code,
		Message : message,
		Data : data,
	}
}

/**
 * @title HttpReturn
 * @Description: http请求返回
 * @param code code值
 * @param message 返回信息
 * @param data 返回数据值
 * @Date 2021-10-20 14:55:39
 **/
func (g *Gin) HttpReturn(code int, message string, data interface{}) {
	g.Ctx.JSON(http.StatusOK, Response{
		Code : code,
		Message : message,
		Data : data,
	})
	return
}
