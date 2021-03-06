package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var aa = 3		//函数外面不可以用:= ,因为在外面需要关键字进行开头（包内部的变量）
var bb = 4
var cc = 5

var (			//整合到一起
	aaa = 33
	bbb = 44
	ccc = 55
)

func variableZeroValue()  {
	var a int
	var s string
	fmt.Printf("%d, %q\n",a, s)
}

func variableInittalValue()  {
	var a, b int = 4, 5
	var s string = "abc";
	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	var a, b, c, d  = 7, 8, false, "deduction"
	fmt.Println(a, b, c, d)
}

func variableShorter()  {
	a, b, c, d  := 7, 8, false, "deduction"		//:= 代表初始化赋值
	b = 10
	fmt.Println(a, b, c, d)
}

func eular()  {
	e := 3 + 4i
	fmt.Println(cmplx.Abs(e))
}

func triangle()  {		//类型都是强制转换的
	var a ,b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts()  {
	const filename  = "abc.txt"	//const 使用与var相似
	const a, b  = 3, 4		//不约定类型则类型不确定[常量相当于文本替换，常量的数值可以作为任何类型使用]
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums()  {
	const (
		cpp = iota  //自增值
		java
		python
		golang
	)

	const (
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInittalValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aaa, bbb, ccc)

	eular()
	triangle()
	consts()
	enums()
}
