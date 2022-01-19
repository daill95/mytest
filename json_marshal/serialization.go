package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Birthday string `json:"birthday"`
	Sal float64 `json:"sal"`
	Skill string `json:"skill"`
}

func testStructJson(){
	monster := Monster{
		Name : "牛魔王",
		Age : 500,
		Birthday : "2000-002-002",
		Sal:  12321.0,
		Skill : "牛魔犬",
	}

	date, err := json.Marshal(&monster)
	if err != nil{
		fmt.Println("序列化失败")
	}
	fmt.Printf("序列化后：%v\n", string(date))
}

func testSliceJson(){
	var slicea []map[string]interface{}
	var ms1 map[string]interface{}
	ms1 = make(map[string]interface{})
	ms1["name"] = "jack"
	ms1["age"] = 7
	ms1["address"] = "北京"

	var ms2 map[string]interface{}
	ms2 = make(map[string]interface{})
	ms2["name"] = "tom"
	ms2["age"] = 17
	ms2["address"] = "华盛顿"

	slicea = append(slicea, ms1)
	slicea = append(slicea, ms2)

	sdata, _ := json.Marshal(slicea)
	fmt.Printf("输出：%v\n", string(sdata))

}

func testMapJson(){
	var minfo map[string]interface{}
	minfo = make(map[string]interface{})
	minfo["name"] = "红孩儿"
	minfo["age"] = 102
	minfo["num"] = 123

	mdata, _:= json.Marshal(minfo)
	fmt.Printf(string(mdata))
	fmt.Println()
}

func main()  {
	testStructJson()
	testMapJson()
	testSliceJson()
}
