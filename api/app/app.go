package app

import (
	"github.com/billcoding/diskplayer/api/web"
	"github.com/gin-gonic/gin"
	"github.com/rwscode/svc"
)

var _app = svc.GetApp(gin.Recovery(), svc.Cors(), svc.WebTry(web.Fs))

func GetApp() *gin.Engine { return _app }
