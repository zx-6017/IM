package middlewares

import (
	"IM/helpers"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"time"
)

func Request() gin.HandlerFunc{

	return func(context *gin.Context) {
		traceId := helpers.GetUUid()
		startTime := time.Now()

		logs.Info("request start : %s %s",traceId,startTime)

		context.Next()

		var(
			now = time.Now().Format("2006-01-02 15:04:05.000")
			duration = int(time.Now().Sub(startTime)/1e6)
			request = context.Request.RequestURI
			host = context.Request.Host
			clientIp = context.ClientIP()
			code = context.Writer.Status()
			ua = context.Request.UserAgent()
		)

		logs.Info("request end : %s %s %d %s %s %s %d %s",traceId,now,duration,request,host,clientIp,code,ua)




	}
}