package main

import (
	"fmt"
	"log"
	"net/http"
)

// 不知为啥，在main函数中直接调用 httpGet() 等函数可以，但包一层(即调用本函数)不行
func httpPrint() {
	httpGet()
}

// 🚀error接口

/*
1. 请求分成两个函数，第一部分如 httpGet()，函数放目录(/gins)和请求地址(localhost:8002)，第二部分如 handler1()，放响应的内容
2. 两个函数不能使用相同的端口，否则第二个调用的不生效
*/
// get 简单请求
func httpGet() {
	http.HandleFunc("/gins", handler1) // 前者定义后缀，后者定义响应
	log.Fatal(http.ListenAndServe("localhost:8011", nil))
}

// 同一端口多个URL的请求 (业务使用)
func httpMultiURL() {
	http.HandleFunc("/gins", handler1) // 多个URL的写法，一个URL(如"/gins")对应一个函数(如handler1))
	http.HandleFunc("/", handler2)
	http.HandleFunc("/list", handler3)
	log.Fatal(http.ListenAndServe("0.0.0.0:8024", nil)) // 注意后面这个参数都填nil(原因省略很多字)
	//log.Fatal(http.ListenAndServe("localhost:8003", nil)) // 注意后面这个参数都填nil(原因省略很多字)
}

// 请求的响应函数
func handler1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World 4043")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nice! URL.Path = %q\n", r.URL.Path)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HeyHey, url = list, handler = 3")
}

/*
启动服务器
$ go run aaaStart.go // 运行
$ go run aaaStart.go & // 后台运行
*/
