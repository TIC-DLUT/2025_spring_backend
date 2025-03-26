package user

import (
	"chatbox/servicecontext"
	"chatbox/tool"
	"chatbox/types"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Telephone string `json:"telephone" bg:"must"`
	Password  string `json:"password" bg:"must"`
}
type LoginResp struct {
	types.BaseResp
	Token string `json:"token"`
}

func NewLoginRoute(sctx *servicecontext.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &LoginReq{}
		e := tool.ParseForm(req, ctx)
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		resp, e := Login(sctx, req)
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		ctx.JSON(200, resp)
	}
}
func Login(sctx *servicecontext.ServiceContext, req *LoginReq) (resp *LoginResp, e error) {

	// JWT详细介绍：
	// https://dinglz.cn/p/jwt%E7%9A%84%E5%BA%94%E7%94%A8/

	// 此处可以做对telephone的校验：
	// 比如是否为手机号的正确格式
	// 增加验证码逻辑防止恶意注册

	// 取到telephone对应的用户信息
	res, e := sctx.UserModel.GetByTelephone(req.Telephone)
	if e != nil {
		// 可能是与数据库链接断开、或者已经有这个telephone等等
		// 精细点可以对telephone存在做一个特判，然后不返回错误，Message置为该telephone已存在
		return
	}
	resp = &LoginResp{
		BaseResp: types.BaseResp{
			Code: 0,
		},
	}
	// 因为注册时放置的是密码的md5加密后的数据，因此对比需要把明文也进行md5加密
	if res.Password == tool.GenerateMD5(req.Password) {
		resp.Message = "登录成功！"
		// 生成JWT
		resp.Token = tool.GenerateJWToken(sctx.Config.JWTPassword, res.ID, res.Telephone)
	} else {
		resp.Code = 1
		resp.Message = "密码错误！"
	}
	return
}
