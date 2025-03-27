package record

import (
	"chatbox/servicecontext"
	"chatbox/tool"
	"chatbox/types"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type DeleteReq struct {
}
type DeleteResp struct {
	types.BaseResp
}

func NewDeleteRoute(sctx *servicecontext.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &DeleteReq{}
		e := tool.ParseForm(req, ctx)
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		resp, e := Delete(sctx, req, ctx.GetString("telephone"))
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		ctx.JSON(200, resp)
	}
}
func Delete(sctx *servicecontext.ServiceContext, req *DeleteReq, telephone string) (resp *DeleteResp, e error) {

	// 获取应用所在的根目录
	rootPath, _ := os.Getwd()
	// 取到存储路径
	workPath := filepath.Join(rootPath, "data", "record")

	// 如果存储的文件存在，删除即可
	if tool.FileExist(filepath.Join(workPath, telephone+".json")) {
		os.Remove(filepath.Join(workPath, telephone+".json"))
	}

	return &DeleteResp{}, nil
}
