package middleware

import (
	"chatbox/servicecontext"
	"chatbox/tool"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func UserAccessMiddleware(sctx *servicecontext.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ""

		// 我们规定Token要么放在Header里，要么放在Cookie里传输
		tokenHeader := ctx.GetHeader("Token")
		tokenCookie, _ := ctx.Cookie("Token")

		// 都不存在，说明未登录
		if tokenHeader == "" && tokenCookie == "" {
			ctx.JSON(401, map[string]interface{}{
				"code":    -1,
				"message": "请先登录！",
			})

			// 暂停后续调用，即不进入真正的功能模块
			ctx.Abort()
			return
		}

		// 取一个有值的Token
		if tokenHeader == "" {
			token = tokenCookie
		} else {
			token = tokenHeader
		}

		// 解析Token
		identify, e := tool.ParseJWToken(sctx.Config.JWTPassword, token)

		// 解析失败：说明过期或者无效
		if e != nil {
			ctx.JSON(401, map[string]interface{}{
				"code":    -1,
				"message": "Token过期！",
			})
			ctx.Abort()
			return
		}

		// 把telephone存到ctx中，方便后续功能
		// 比如record/add根据telephone存储信息
		ctx.Set("telephone", identify.Telephone)

		// 继续下一个模块
		ctx.Next()
	}
}
