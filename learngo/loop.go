package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

/**
	十进制转换成2进制
 */
func convertToBin(n int) string {
	result := ""
	for ; n>0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string)  {
	file, err := os.Open(filename)		//打开文件获得句柄

	if  err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(file)		//一行一行的读取

	for scanner.Scan()  {					//循环偏移（相当于while）
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(7212612712),
	)

	printFile(".gitignore")
}
