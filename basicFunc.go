package main

import "fmt" // Go 中要打印一个值需要引入 fmt 包

func basicFuncPrint() {
	// print println printf
	println("basicFunc: ➤ ")
	// println 换行、多值间有空格
	fmt.Println("println", "go", "python", "php", "javascript") // go python php javascript
	// print 不换行、多值间无空格
	fmt.Print("print", "go", "python", "php", "javascript", "\n") // gopythonphpjavascript
	const (
		ta = iota
		tb
		tc
	)
	// printf 格式化打印(%d等)，不换行
	fmt.Printf("printf: %d, tb = %d, tc = %d end", ta, tb, tc)

	// len
	var a = [3]int{1, 2, 3}
	fmt.Println("len", len(a))

}
