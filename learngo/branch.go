package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string {
	g:= ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d", score))    	//中断程序的运行(Sprint是返回填写后的结果，而不是将结果写入 os.Stdout 中)
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}

	return g
}

func main() {
	const filename  = ".gitignore"
	contents, err  := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}
	//等效于
	if con, errs  := ioutil.ReadFile(filename) ;errs != nil {	//此时con的作用域仅在if中
		fmt.Println(errs)
	}else {
		fmt.Printf("%s\n", con)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		grade(101),
	)
}
