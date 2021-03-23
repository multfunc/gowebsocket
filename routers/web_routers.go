/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:20
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"gowebsocket/controllers/home"
	"gowebsocket/controllers/systems"
	"gowebsocket/controllers/user"
)

func Init(router *gin.Engine) {
	router.LoadHTMLGlob("views/**/*")

	// 用户组
	userRouter := router.Group("/im/user")
	{
		userRouter.GET("/list", user.List)
		userRouter.GET("/online", user.Online)
		userRouter.POST("/sendMessage", user.SendMessage)
		userRouter.POST("/sendMessageAll", user.SendMessageAll)
	}

	// 系统
	systemRouter := router.Group("/im/system")
	{
		systemRouter.GET("/state", systems.Status)
	}

	// home
	homeRouter := router.Group("/im/home")
	{
		homeRouter.GET("/index", home.Index)
	}

	// router.POST("/user/online", user.Online)
}
