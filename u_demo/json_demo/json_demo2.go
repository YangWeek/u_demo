package jsondemo

import (
	"encoding/json"
	"fmt"
	"math"
)

type MyData struct {
	ID   int64  `json:"id,string"` //  JSON 序列化和反序列化的结构体标签。它的作用是指定在 JSON 数据中字段的名称和类型
	Name string `json:"name"`
}

//func (d *MyData) Unmarshal() {
//
//}
//
//func (d *MyData) Marshal() {
//
//}

// 第一层：
// 第二层：
// 第五层：
// 第九层：

func Init_jsondemo2_test2() {
	// 序列化： 后端的数据->JSON格式的数据
	d1 := MyData{
		ID:   math.MaxInt64,
		Name: "七米",
	}
	// json序列化
	b, err := json.Marshal(d1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// 反序列化：JSON格式的数据 -> Go语言中的数据
	s := `{"id":"9223372036854775807","name":"七米"}`
	var d2 MyData
	// 强制转化为 []byte 类型
	if err := json.Unmarshal([]byte(s), &d2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v type:%T\n", d2, d2.ID)
}
