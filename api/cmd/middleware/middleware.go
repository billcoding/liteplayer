package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rwscode/svc"
)

func init() { svc.GetApp().Use(gin.Recovery(), Cors()) }
