package routers

import (
	"IM/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine){

	api := router.Group("/api")
	{
		api.GET("/create",controllers.Create);
	}


}
