package main

import (
	"encoding/json" // json库
	"fmt"
	"log"
)

// !+
type Movie struct {
	Title  string // 只有可导出变量可以转换为JSON字段，所以变量都是首字母大写
	Year   int    `json:"released"`        // `json:"released"` 成员标签定义(field tag)，格式为`key:"value"`。此处实现输出到json时，重命名为"released"
	Color  bool   `json:"color,omitempty"` // omitempty 表示如果成员的值是零值或空，则不输出到json中
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

//!-

func jsonPrint() {
	{
		// Marshal：Go数据结构格式化为json
		data, err := json.Marshal(movies) // Marshal：把movies转为json格式
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("Marshal：数组格式化为json：%s\n", data)
	}

	{
		//MarshalIndent：Go数据结构格式化为json+带缩进
		data, err := json.MarshalIndent(movies, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("MarshalIndent：数组格式化为json+带缩进：%s\n", data)

		//Unmarshal：json解码为Go数据结构，逆操作
		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println("Unmarshal：json解码为Go数据结构，逆操作：", titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
		//!-Unmarshal
	}
}

/*
//!+output
[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingr
id Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Ac
tors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"
Actors":["Steve McQueen","Jacqueline Bisset"]}]
//!-output
*/

/*
//!+indented
[
    {
        "Title": "Casablanca",
        "released": 1942,
        "Actors": [
            "Humphrey Bogart",
            "Ingrid Bergman"
        ]
    },
    {
        "Title": "Cool Hand Luke",
        "released": 1967,
        "color": true,
        "Actors": [
            "Paul Newman"
        ]
    },
    {
        "Title": "Bullitt",
        "released": 1968,
        "color": true,
        "Actors": [
            "Steve McQueen",
            "Jacqueline Bisset"
        ]
    }
]
//!-indented
*/
