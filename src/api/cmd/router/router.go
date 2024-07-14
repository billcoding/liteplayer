package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rwscode/svc"
)

func init() {
	svc.GetApp().GET("/api/echo", echo)
}

func echo(ctx *gin.Context) {
	type resp0 struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	}
	svc.QueryResp(ctx, func() (resp resp0, err error) { return resp0{1000, "name"}, nil })
}
