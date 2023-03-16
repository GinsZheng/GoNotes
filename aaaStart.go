package main // 几乎每个文件都需要

import (
	"fmt"
	"log"
	"net/http"
) // 导入

func main() { // main 函数只能有一个，是个特殊的已定义函数
	fmt.Println("Hello World") // 打印

	//osPrint()

	//getDomain()
	//slicePrint()

	//mysqlPrint()
	//uploadImage()

	//uploadImageHTML()
	//importPrint()
	//httpPrint()

	//testPrint()                 // 在主函数里可以执行同一文件的其他函数
	//declarationPrint()          // 在主函数里执行同一目录下其他文件(declaration.go)的函数
	//ArrayPrint()
	//slicePrint()
	//basicFuncPrint()
	//serverTest()
	//structPrint()
	//jsonPrint()
	//funcPrint()

	httpMultiURL()
	//httpGet() // 关于http型的函数，似乎只会执行第一个(可能要用并发方式来执行后面的)

}

func testPrint() {
	fmt.Println("Test1")
}

func serverTest() {
	http.HandleFunc("/", handler) // each request calls handler
	//log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// HTTP请求的一个参数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nice! URL.Path = %q\n", r.URL.Path)
}
