package main

// 每一个Go源文件的开头都要进行包声明，目的是当该包被其他包引入时作为其默认的标识符(称为包名)
// 如math/rand包中每一个文件的开头都是 package rand，这样当你导入这个包时，可以访问它的成员，如rand.int、rand.Float64等
// 包名不包含后缀(如在导入路径发问追加版本号后缀"gopkg.in/yaml.v2"，此时包名为"yaml")

import ( // import 多个时，可用括号
	"crypto/rand"
	"fmt" // 可以用空行↓来分组

	mrand "math/rand" // 如果两个包名相同，要把其中一个重命名，称为重命名导入(如此处重命名为mrand，mrand只在当前文件生效)
	// 重命名导入在没有冲突是也非常有用(特别是导入的包名很长时)

	_ "image/png" // 空导入。背景：如果一个包未在文件中用到，那么保存时会自动清除。而有的时候我们必须导入一个包，仅仅是为了利用其副作用：对包级别的变量执行初始化表达式求值，并执行它的init函数。此时为了避免被清除，就在包名前加一个⌜_⌟来实现。本例为注册PNG解码器，没有此空导入，程序可以正常编译和链接，但不能识别和解码PNG格式的输入
)

func importPrint() {
	fmt.Println("importTest")
	fmt.Println(rand.Reader)
	fmt.Println(mrand.Int())

}
