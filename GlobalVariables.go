package main

import (
	"fmt"
	"runtime"
)

func osPrint() {
	fmt.Println("OS Type is:", runtime.GOOS) // 判断当前运行操作系统类型。MacOS为：darwin ，Centos为：linux
}

// 全局变量
const testIP = "127.0.0.1"
const productionIP = "0.0.0.0"
const port = "8008"
const domain = "ginkgeek.com"
const testAddress = testIP + ":" + port
const productionAddress = productionIP + ":" + port

func getDomain() string {
	if runtime.GOOS == "darwin" {
		return testAddress
	} else {
		return productionAddress
	}
}
