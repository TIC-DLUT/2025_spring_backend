package ai

import (
	"chatbox/servicecontext"
	"chatbox/tool"
	"encoding/json"

	"github.com/dingdinglz/openai"
	"github.com/gin-gonic/gin"
)

type RunReq struct {
	Question string `json:"question" bg:"must"`
}

func NewRunRoute(sctx *servicecontext.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 该接口比较特殊，需要以SSE的方式进行流式传输

		// 设置头
		// 标注特殊类型：SSE
		ctx.Header(`Content-Type`, `text/event-stream`)
		// 要求禁止缓存，否则可能导致实现不了流式效果
		ctx.Header("Cache-Control", "no-cache")
		// 保持长链接
		ctx.Header("Connection", "keep-alive")
		// 转发也禁止缓存
		ctx.Header("X-Accel-Buffering", "no")
		req := &RunReq{}
		e := tool.ParseQuery(req, ctx)
		if e != nil {
			ctx.SSEvent("message", e.Error())
			return
		}

		// openai包的使用方法见：
		// https://github.com/dingdinglz/openai
		client := openai.NewClient(&openai.ClientConfig{
			BaseUrl: sctx.Config.BasePath,
			ApiKey:  sctx.Config.ApiKey,
		})
		e = client.ChatStream(sctx.Config.Model, []openai.Message{
			{Role: "system", Content: "你是一个智能问答机器人"},
			{Role: "user", Content: req.Question},
		}, func(s string) {
			// 首先用JSON打包防止换行等特殊字符取不到
			message := map[string]string{
				"data": s,
			}
			messageJSON, _ := json.Marshal(message)
			// 传输一次
			ctx.SSEvent("message", string(messageJSON))
			ctx.Writer.Flush()
		})
		if e != nil {
			ctx.SSEvent("message", e.Error())
		}
	}
}
