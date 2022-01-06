package rest

import "github.com/gin-gonic/gin"

type Response struct{
	c *gin.Context
}

func (r *Response) Response(httpCode int, errCode string, data interface{}) *Response{
	r.c.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg": errCode,
		"data": data,
	})
	return r
}

