package record

import (
	"chatbox/servicecontext"
	"chatbox/tool"
	"chatbox/types"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type getReq struct {
}
type getResp struct {
	types.BaseResp
	Data []string `json:"data"`
}

func NewGetRoute(sctx *servicecontext.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &getReq{}
		e := tool.ParseQuery(req, ctx)
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		resp, e := get(sctx, req, ctx.GetString("telephone"))
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		ctx.JSON(200, resp)
	}
}
func get(sctx *servicecontext.ServiceContext, req *getReq, telephone string) (resp *getResp, e error) {

	// 获取应用所在的根目录
	rootPath, _ := os.Getwd()
	// 取到存储路径
	workPath := filepath.Join(rootPath, "data", "record")

	var records []string

	// 跟add一样，如果文件存在，就读入records
	// 不存在留空也无所谓，本来不存在数组就要置空
	if tool.FileExist(filepath.Join(workPath, telephone+".json")) {
		var recordBytes []byte
		recordBytes, e = os.ReadFile(filepath.Join(workPath, telephone+".json"))
		if e != nil {
			return
		}
		e = json.Unmarshal(recordBytes, &records)
		if e != nil {
			return
		}
	}
	return &getResp{
		Data: records,
	}, nil
}
