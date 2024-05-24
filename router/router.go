package router

import (
	"GoChat/middlewear"
	"GoChat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")

	// 用户模块
	user := v1.Group("user")
	{
		user.GET("/list", middlewear.JWY(), service.List)
		user.POST("/login_pw", service.LoginByNameAndPassWord)
		user.POST("/new", service.NewUser)
		user.DELETE("/delete", middlewear.JWY(), service.DeleteUser)
		user.POST("/updata", middlewear.JWY(), service.UpdataUser)
	}

	// 关系模块
	reletion := v1.Group("relation").Use(middlewear.JWY())
	{
		reletion.POST("/list", service.FriendList)
		reletion.POST("/add", service.AddFriendByName)
		reletion.POST("/new_group", service.NewGroup)
		reletion.POST("/group_list", service.GroupList)
		reletion.POST("/join_group", service.JoinGroup)
	}

	return router
}
