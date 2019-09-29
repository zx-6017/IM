package routers

import (
	"IM/controllers"
	"IM/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine){

	api := router.Group("/api").Use(middlewares.Request())
	{
		api.GET("/create",controllers.Create)
		api.GET("/friendrelation",controllers.FriendRelation)
	}


}
