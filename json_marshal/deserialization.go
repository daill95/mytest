package main

import (
	"encoding/json"
	"fmt"
)

//反序列化成一个Monster结构体类型
func unMarshalStruct() {
	str1 := "{\"name\":\"牛魔王\",\"age\":500,\"birthday\":\"2000-002-002\",\"sal\":12321,\"skill\":\"牛魔犬\"}"

	type Monster struct {
		Name     string  `json:"name"`
		Age      int     `json:"age"`
		Birthday string  `json:"birthday"`
		Sal      float64 `json:"sal"`
		Skill    string  `json:"skill"`
	}
  
	var monster Monster
	//这里要传&monster，传地址，因为要在函数中修改monster
	err := json.Unmarshal([]byte(str1), &monster)
	if err != nil{
		fmt.Println("反序列号失败")
	}
	fmt.Println(monster.Name)
}
func main()  {
	unMarshalStruct()
}
