## 什么是api

[什么是API](https://zh.wikipedia.org/wiki/%E5%BA%94%E7%94%A8%E7%A8%8B%E5%BA%8F%E6%8E%A5%E5%8F%A3)（来自维基百科）

然而今天我们要谈论的并不是这种广义上的api，而是web api，这种一般是服务提供商方便我们使用三方服务提供的接口，一般通过http去调用，也有类似提供websocket或者sse的接口（参见这次vivo aigc提供的接口）

首先，我们以一个简单的接口举例：[一言接口文档](https://developer.hitokoto.cn/sentence/)

这个接口之前例会拿cpp调用演示了一次，这次我们用go来调用，并研究更深层次的用法。

首先拿到一个接口文档，我们要关注一下内容：

1. 调用方式（GET还是POST还是其他？）
2. api地址，比如本接口中的：`v1.hitokoto.cn`
3. 请求参数：
![](image/4/1.png)

4. 返回参数：
![](image/4/2.png)
传入参数的方式，GET一般是query，也就是?name=&pass=这样的格式，POST一般是form。

对应到gin的解析方式，一个是`ctx.Query()`，一个是`ctx.PostForm()`

当然，有的时候POST也可以要求用query传数据。

至于返回数据，目前最主流的方式就是JSON，之前例会我们也介绍过了。

## 调用一言接口

首先它是一个GET接口，让我们先不带任何参数，请求一下，看看返回的格式是如何的。

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, e := http.Get("https://v1.hitokoto.cn/")
	if e != nil {
		panic(e)
	}
	if res.StatusCode != http.StatusOK {
		panic("请求失败！")
	}
	defer res.Body.Close() // 保证在程序结束关闭连接
	body, e := io.ReadAll(res.Body)
	if e != nil {
		panic(e)
	}
	bodyText := string(body)
	fmt.Println(bodyText)
}
```

可以看到，返回的格式如下：

```json
{
    "id": 479,
    "uuid": "05f790b7-527b-4aa3-9112-365e061cfa25",
    "hitokoto": "透过孩子的眼神，让我相信这个世界上还有着纯真。",
    "type": "g",
    "from": "宫崎骏",
    "from_who": null,
    "creator": "anythink",
    "creator_uid": 0,
    "reviewer": 0,
    "commit_from": "web",
    "created_at": "1468950861",
    "length": 23
}
```

要取回数据，只需要解析返回的json就可以了。

### Go解析json的方法

golang提供了`encoding/json`包实现序列化和反序列化json，想要反序列化json，（即把json变成我们可以操作的对象）我们通常是把一个json绑定到一个对应的结构体，比如：

```json
{
	"name": "dinglz",
	"age": 18
}
```

可以解析到如下结构体

```go
type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
```

这里，由于结构体不允许公开参数首字母是小写，所以我们需要用一个json tag去把json的字段对应到结构体的字段，然后可以用`json.Unmarshal`去把json的数据解析到struct上

```go
data := Person{}
json.Unmarshal([]byte(jsonText) , &data)
```

除了解析到结构体，也可以解析到`map[string]interface{}`，之前说过，`interface{}`也就是`any`，可以灵活的转换成任意结构，这也是go在设计中对泛型的妥协（虽然现在已经支持了泛型的写法，interface{}仍然是很通用的写法）

---

那么我们这里先用结构体的方法对结果进行解析，这里给大家一个网站，可以根据一个json的格式生成对应的go struct：[点我使用](https://www.bejson.com/transfor/json2go/)

> [!NOTE]
> 不需要工具我们当然也可以手写出用来存储数据的结构体，但是开发效率和代码质量是衡量一个优秀的程序员的标准，当我们巧妙的使用工具，提升自己的开发效率，才是正确的选择。
> AI同理，我们当然可以合理的利用它，但当你不知道它在干什么时，最好不用使用它，因为我们应该是让AI帮我们写已经会的内容，或者启发我们的思路，而不是看不懂AI的代码就直接用，这样无法让你的编码能力得到任何的提升。

那么对应的结构体如下：

```go
type Data struct {
	ID int `json:"id"`
	UUID string `json:"uuid"`
	Hitokoto string `json:"hitokoto"`
	Type string `json:"type"`
	From string `json:"from"`
	FromWho interface{} `json:"from_who"`
	Creator string `json:"creator"`
	CreatorUID int `json:"creator_uid"`
	Reviewer int `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt string `json:"created_at"`
	Length int `json:"length"`
}
```

那么我们解析一下，并把一言的内容输出出来：

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	ID         int         `json:"id"`
	UUID       string      `json:"uuid"`
	Hitokoto   string      `json:"hitokoto"`
	Type       string      `json:"type"`
	From       string      `json:"from"`
	FromWho    interface{} `json:"from_who"`
	Creator    string      `json:"creator"`
	CreatorUID int         `json:"creator_uid"`
	Reviewer   int         `json:"reviewer"`
	CommitFrom string      `json:"commit_from"`
	CreatedAt  string      `json:"created_at"`
	Length     int         `json:"length"`
}

func main() {
	res, e := http.Get("https://v1.hitokoto.cn/")
	if e != nil {
		panic(e)
	}
	if res.StatusCode != http.StatusOK {
		panic("请求失败！")
	}
	defer res.Body.Close() // 保证在程序结束关闭连接
	body, e := io.ReadAll(res.Body)
	if e != nil {
		panic(e)
	}
	data := Data{}
	e = json.Unmarshal(body, &data)
	if e != nil {
		panic(e)
	}
	fmt.Println(data.Hitokoto)
}
```

这下运行程序就能直接输出一句话了，如果我们想要输出其他信息，比如作者等等，通过`Data.From`结构体的参数也能轻松取到。

之前的例会我们到此为止了，下面让我们继续探索一下请求参数：比如我们只想要诗词的内容，根据文档的描述，我们就需要指定c为i，那么让我们修改URL加上query：`?c=i`即可。（如果有多个query，可以用&链接，比如`?c=i&encode=json`）

> 需要注意，当query中含有汉字或者特殊字符时，需要进行url编码，有很多包支持类似的功能，不再赘述

加上类型指定后的代码：

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	ID         int         `json:"id"`
	UUID       string      `json:"uuid"`
	Hitokoto   string      `json:"hitokoto"`
	Type       string      `json:"type"`
	From       string      `json:"from"`
	FromWho    interface{} `json:"from_who"`
	Creator    string      `json:"creator"`
	CreatorUID int         `json:"creator_uid"`
	Reviewer   int         `json:"reviewer"`
	CommitFrom string      `json:"commit_from"`
	CreatedAt  string      `json:"created_at"`
	Length     int         `json:"length"`
}

func main() {
	res, e := http.Get("https://v1.hitokoto.cn/?c=i")
	if e != nil {
		panic(e)
	}
	if res.StatusCode != http.StatusOK {
		panic("请求失败！")
	}
	defer res.Body.Close() // 保证在程序结束关闭连接
	body, e := io.ReadAll(res.Body)
	if e != nil {
		panic(e)
	}
	data := Data{}
	e = json.Unmarshal(body, &data)
	if e != nil {
		panic(e)
	}
	fmt.Println(data.Hitokoto)
}
```

这下取出的内容只有诗词了。

当然，不同的api有不同的文档，上面只是展示了一言api的使用方法，调用其他的api请以文档为准。

## http client包

我们刚刚用的是官方的http请求包，功能可能较少，给大家推荐一个非常好用的client包，功能很多：[Resty](https://resty.dev/)

## AI LLM 接口

接下来的项目中我们需要调用ai服务商的api，这些涉及到sse等可能比较复杂，我们在项目中会使用我封装好的包：[OpenAI](https://github.com/dingdinglz/openai)，这个包最初是为了方便软创项目中对AI的调用封装而成，最后被我剥离出了做成了个三方包。

readme写的相对来说比较详细，想了解用法可以直接看readme

### AI LLM接口的综合利用

大家平常生活中应该也都用过ai大模型对话，像deepseek啊，kimi啊，豆包啊。

那么在我们开发中，如何利用ai大模型的接口呢，难道功能就局限于聊天嘛？其实并不是。

在常理来看，ai返回的内容是没有具体的结构约束的，因此，很难做返回数据的处理，也就是很难把ai当成一个工具来用。

但，让我们想想我们能解析的结构有哪些：我们之前提到过的json，一些其他的数据格式像yml，这些是可以解析到的，因此我们只需要在对话的提示词中做出限制就可以让ai成为一个工具包。

例如，我们想让ai计算一个数值，正常比如可能问1+1等于多少，不同家的ai回答的结果都有可能不同，尽管答案可能都是1。

```
DeepSeek: 1 + 1 等于 **2**。

这是最基本的数学加法，如果有其他上下文或特殊定义，结果可能会不同（比如在某些二进制或逻辑运算中 1 + 1 = 10），但在常规算术中，答案就是 **2**。 😊

Kimi：1+1等于2。

通义千问：1 + 1 等于 **2**。

因此，答案是：**2​**
```

很显然，结果都是对的，但这是我们想要的结果吗，或许不是，我们怎么取到2呢？（要注意，每次问的回答可能都不一样，因此对于这样的回答，很难取到我们想要的结果）

但是，提示词可以帮助我们限制回答的格式，比如在问之前，加上`下面需要你计算一个算式。返回的回答以json为格式，例如，如果答案是1，就返回{"value":1}，如果是其他答案就替换value的值，回答只需要出现json的内容即可，不需要包括markdown代码框在内的任何内容`

![](image/4/3.png)

这下

```json
{
	"value" : 2
}
```

很显然就是我们可以解析的格式了

以上只是举个例子，帮助你明白调用ai llm接口的一些小技巧。

json可以解决很多情况，但是消耗token数较大，因此大家现在往往不采用json去解决问题，而是使用yml，相对较为省token。

还有种情况：流式传输时（一点一点返回结果）

json很显然是无法被解析的（比如只返回了`{"value":` ），但是yml可以做到，或者用其他一些特殊的办法比如自己设计格式，ans1,ans2,ans3...这种，这些都是需要通过对提示词来做手脚做到的

利用提示词的同时也要注意提示词被注入问题。我们看一个例子大家应该就能理解了，也就是前段时间比较出名了的小红书翻译（利用llm进行翻译，我猜提示词可能是翻译下面的文字，只需要给出翻译的内容即可）：

![](image/4/4.png)

可以看到，括号里的内容就是要求llm在完成提示词的任务后，再进行额外的工作，大模型欣然听命了：

![](image/4/5.png)

因此，我们要注意保证下提示词的健壮性（有点难哈），或者在传入前就审核一下传入的内容（比如多用一次ai llm，这次提示词就让他审核有无提示词注入的部分）

简单讲到这吧，后面比赛大家会用到很多llm api的内容的，到时候欢迎大家来交流使用心得。

---

[回到主页](README.md)