package main

import (
	"chatbox/config"
	"chatbox/database"
	"chatbox/servicecontext"
	"chatbox/tool"
	"flag"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	etc string
)

func init() {
	flag.StringVar(&etc, "etc", "config.json", "config file path")
	flag.Parse()
}

func main() {
	// 创建存储目录
	rootPath, _ := os.Getwd()
	tool.CreateDir(filepath.Join(rootPath, "data", "record"))

	config := config.Load(etc)
	database.Open("sqlite", "data.db")
	scvx := servicecontext.NewServiceContext(config)
	server := gin.New()
	BindRoute(scvx, server)
	e := server.Run(config.Host + ":" + strconv.Itoa(config.Port))
	if e != nil {
		panic(e)
	}
}
