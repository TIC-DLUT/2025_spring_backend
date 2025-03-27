package main

import (
	"chatbox/middleware"
	"chatbox/route/record"
	"chatbox/route/user"
	"chatbox/servicecontext"

	"github.com/gin-gonic/gin"
)

func BindRoute(scvx *servicecontext.ServiceContext, server *gin.Engine) {

	serverusergroup := server.Group("/user")
	serverusergroup.POST("/register", user.NewRegisterRoute(scvx))
	serverusergroup.POST("/login", user.NewLoginRoute(scvx))
	serverusergroup.POST("/check", user.NewCheckRoute(scvx))

	serverrecordgroup := server.Group("/record", middleware.UserAccessMiddleware(scvx))
	serverrecordgroup.POST("/add", record.NewAddRoute(scvx))
	serverrecordgroup.GET("/get", record.NewGetRoute(scvx))
	serverrecordgroup.POST("/delete", record.NewDeleteRoute(scvx))

}
