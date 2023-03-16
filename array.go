package main

import "fmt"

/*
数组基本知识：
Go中的数组长度固定，数组很少用，一般用Slice
*/
func ArrayPrint() {
	// 3个整数的数组：[0,0,0]
	var a [3]int
	fmt.Println("ArrayPrint:", a, "a[0]:", a[0], "len:", len(a))

	// 3个特定内容的数组
	var b = [3]int{1, 2, 3} //
	fmt.Println("b =", b)

	// n个特定内容的数组
	var c = [...]int{1, 2, 3, 4}
	fmt.Println("c =", c)

	// 定义n个0
	var d = [...]int{1, 10: -1, -2} // 10后面加冒号，表示10个0
	fmt.Println("d", d)
}

/*
slice是没有固定长度的数组
两个slice不允许直接比较，唯一允许的比较操作是和nil比较，但注意，如果想智能slice是否为空，要用len(s) = 0，而不是s == nil，因为可能s为空时，不为nil(比如赋值为 s = []int{} 时)
*/
func slicePrint() {
	fmt.Println("slicePrint ➤ ")
	// 创建一个特定slice。[]int 中括号中不带东西，就是切片类型
	var a = []int{1, 2, 3}
	fmt.Println("a", a)
	a = append(a, 4) // append 这么写
	fmt.Println("a", a)

	// 创建一个空切片
	var s []string
	s = append(s, "a")
	fmt.Println("s", s)
}
