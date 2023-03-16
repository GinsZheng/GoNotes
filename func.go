package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func funcPrint() {
	fmt.Println("基本函数", hypot(3, 4))
	fmt.Println("变量类型简写", f(1, 2, 3, "a", "b"))
	fmt.Println("延迟函数：", trace("test"))
	var p = Point{1, 2}
	var q = Point{4, 6}
	fmt.Println("函数与方法：函数", Distance(p, q))
	fmt.Println("函数与方法：方法", p.Distance(q))
}

func hypot(x float64, y float64) float64 { // 返回值是 float64 (括号外那个)
	return math.Sqrt(x*x + y*y) // 勾股定理求斜边
}

// 变量类型简写
func f(i, j, k int, s, t string) string { // i j k 是同一类型时，可只写一个 int 。s, t 也是
	return s + t
}

// 多返回值函数
//func findLinks(url string) ([]string, error) {
//	return
//}

// 延迟函数：在return之后执行，且可以更新函数的结果变量
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

// Go没有默认参数值

// 方法与函数
// 方法为变量掉用的函数，表现如 x.len()  函数为通用格式，表现如 len(x)

// 辅助项
type Point struct {
	x, y float64
}

// 方法声明
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
	/*
		(p Point) 方法声明相比于普通声明的区别。相当于Swift中的Extension。p 相当于self，可自定义，一般用类型首字母。Point 即要Extenson的类型
	*/
}

// 普通函数声明，与上方实现同样作用
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}
