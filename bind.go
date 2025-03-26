package main

import (
	"chatbox/route/user"
	"chatbox/servicecontext"

	"github.com/gin-gonic/gin"
)

func BindRoute(scvx *servicecontext.ServiceContext, server *gin.Engine) {

	serverusergroup := server.Group("/user")
	serverusergroup.POST("/register", user.NewRegisterRoute(scvx))
	serverusergroup.POST("/login", user.NewLoginRoute(scvx))
	serverusergroup.POST("/check", user.NewCheckRoute(scvx))

}
