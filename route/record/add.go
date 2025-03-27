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

type addReq struct {
	Question string `json:"question" bg:"must"`
}
type addResp struct {
	types.BaseResp
}

func NewAddRoute(sctx *servicecontext.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &addReq{}
		e := tool.ParseForm(req, ctx)
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		resp, e := add(sctx, req, ctx.GetString("telephone"))
		if e != nil {
			ctx.JSON(400, map[string]string{"message": e.Error()})
			return
		}
		ctx.JSON(200, resp)
	}
}
func add(sctx *servicecontext.ServiceContext, req *addReq, telephone string) (resp *addResp, e error) {

	// 获取应用所在的根目录
	rootPath, _ := os.Getwd()
	// 取到存储路径
	workPath := filepath.Join(rootPath, "data", "record")

	var records []string

	// 如果文件存在，说明有历史记录，读取并解析到数组
	// 如果文件不存在，不需要解析，让records留空即可
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

	// 把最新的问题添加到最前面
	// 如果用 records = append(records , req.Question)
	// Question默认会被添加到records的最后
	records = append([]string{req.Question}, records...)

	// 限制record最大只允许有100条，保证存储空间够用，并不保留过早的数据
	if len(records) > 100 {
		records = records[:100]
	}

	// 生成保存文件并保存
	recordJSON, _ := json.Marshal(records)
	e = os.WriteFile(filepath.Join(workPath, telephone+".json"), recordJSON, os.ModePerm)
	resp = &addResp{}
	return
}
