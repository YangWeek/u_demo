package jsondemo

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
)

// Go实战技巧--你不知道的JSON操作技巧
// 当后端传输比较大的json数据  比较大的数字  

// 在后端使用的时候 还是用int64 进行json序列化 和 反序列化  用 string
type Data struct {
	id int64 `json:"id,string"`  // 后端 自动从int64 变成了 string 解决了数字不精确 
	// string：指定 JSON 中 id 的值应被解析为字符串，即使它在 JSON 中是一个数字。
	//Go 的 json 包将会把 JSON 字符串转换为 int64 类型。这种方式可以处理 JSON 数据中数字值的字符串表示
}
func Init_jsondemo_test() {
	r := gin.Default()
	r.LoadHTMLFiles("./json_demo/index.html")
	r.GET("/index",func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK,"index.html",nil)
	})
	r.GET("/data",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,Data{123})
	})

	
	// json.Marshal() 
	r.Run(":8990")
}