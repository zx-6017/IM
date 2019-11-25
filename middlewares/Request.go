package middlewares

import (
	"IM/config"
	"IM/helpers"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strings"
	"time"
)

func Request() gin.HandlerFunc{

	return func(context *gin.Context) {
		traceId := helpers.GetUUid()
		startTime := time.Now()

		logs.Info("request start : %s %s",traceId,startTime)

		defer errEmail(context)

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

func errEmail(request *gin.Context){
	err := recover()
	if err ==nil{
		return
	}

	DebugStack := ""

	for _,v := range strings.Split(string(debug.Stack()),"\n"){
		DebugStack += v
	}
	subject := "【重要错误】项目出错!!! "
	body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
	body  = strings.ReplaceAll(body, "{RequestTime}", "Requesttime")

	body  = strings.ReplaceAll(body, "{RequestURL}", request.Request.Method + "  " + request.Request.Host + request.Request.RequestURI)
	body  = strings.ReplaceAll(body, "{RequestUA}", request.Request.UserAgent())
	body  = strings.ReplaceAll(body, "{RequestIP}", request.ClientIP())
	body  = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

	helpers.SendMail(config.Conf.Strings("MAIL_RECEIVE"),subject,body)

}