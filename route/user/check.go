package user

import (
	"chatbox/servicecontext"
	"chatbox/tool"
	"chatbox/types"

	"github.com/gin-gonic/gin"
)

type CheckReq struct {
	Token string `json:"token" bg:"must"`
}
type CheckResp struct {
	types.BaseResp
	Ok bool `json:"ok"`
}

func NewCheckRoute(sctx *servicecontext.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &CheckReq{}
		e := tool.ParseForm(req, ctx)
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		resp, e := Check(sctx, req)
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		ctx.JSON(200, resp)
	}
}
func Check(sctx *servicecontext.ServiceContext, req *CheckReq) (resp *CheckResp, e error) {
	// 尝试解析token，如果报错说明token失效
	_, err := tool.ParseJWToken(sctx.Config.JWTPassword, req.Token)
	if err != nil {
		return &CheckResp{
			BaseResp: types.BaseResp{
				Code:    1,
				Message: "Token已失效",
			},
			Ok: false,
		}, nil
	}
	// 这里不用设置BaseResp的原因是因为Go的结构体不进行赋值时默认零值
	// 所以code = 0 ， message = ""
	return &CheckResp{
		Ok: true,
	}, nil
}
