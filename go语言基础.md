推荐基础语法学习书籍：[Go入门指南](https://github.com/unknwon/the-way-to-go_ZH_CN)

基础语法的完整学习请参照上述书籍，本文档仅针对可能出错的点，或者需要注意的点进行说明。

## 安装

安装完成后，命令行输入
``` bash
go version
```
如果有版本号输出，说明安装成功，如果显示not found字样，可能是由于未添加到$PATH导致的，自行寻找一下添加到path的方法

### linux添加path

在~/.bashrc添加`export PATH=$PATH:go路径`（针对bash，其他终端可能有不同的配置文件）

### 设置代理

> 注意，安装成功后，代理一定要进行设置，代理用于国内go mod直接访问代码仓库

``` bash
go env -w GOPROXY=https://goproxy.cn,direct
```

## IDE

适合go语言开发的ide推荐两个，一个是vscode安装golang插件，另一个是goland（收费），前者相对轻量级，后者功能更全

### vscode

Go插件安装后会提醒你安装相关工具，如果不设置代理该步就会失败，可以通过命令面板：Go: install/update tools去重新安装工具

![](./image/0/1.png)

## 错误处理

值得一提的是go独特的错误处理机制，在每个函数后几乎都会有错误值的返回，也就是error类型，对于错误的处理应该做到每个错误都进行，而不是为了图省事用_替代

### panic && recover

panic是一个很独特的机制，它将直接使程序崩溃，panic一般用于不可恢复的错误场景，但在服务器开发时，为了保障服务器持续运行，应当在发生panic的情况下仍然保持运行，这个时候需要用到recover方法去做从崩溃中恢复。

各个服务器框架基本都有自己的recover中间件，大家可以通过阅读这些中间件的源码深入了解recover机制的工作原理
- [Fiber recover 中间件](https://github.com/gofiber/fiber/tree/main/middleware/recover)
- [Gin recovery中间件](https://github.com/gin-gonic/gin/blob/master/recovery.go#L33)

## Package

