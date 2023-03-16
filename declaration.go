package main

import "fmt"

var x = 1
var (
	y = 2
	z = 3
)

const a = 5 // const 常量，相当于Swift的 let
const (
	b = 6
	c = 7
)
const (
	Sun = iota // iota 是从0开始的序列
	Mon        // iota之下，后续的内容无需显式定义即可自动+1
	Tue
	Wed
	Thu
	Fri
	Sat
)

func declarationPrint() {
	fmt.Println("x", x, "y", y, "z", z)
	fmt.Println("a", a, "b", b, "c", c)
}
