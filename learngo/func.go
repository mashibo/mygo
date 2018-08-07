package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int , op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q , _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation %s", op)
	}
}

/**
	带余数的除法(函数可以返回多个值)
 */
func div(a, b int) (q, r int) {
	return a / b, a % b
}

/**
	函数式编程（变量传入的是函数）
 */
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	fmt.Println(reflect.ValueOf(op))
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling Function %s with args" + "(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数列表
func sum(numbers ... int) int {
	s := 0
	for i := range numbers{
		s += numbers[i]
	}

	return s
}

func swap(a, b *int)  {		//函数参数只能进行值传递
	*a, *b  = *b, *a
}

func main() {
	if result, error := eval(3, 4, ">"); error != nil {
		fmt.Println("ERROR", error)
	}else {
		fmt.Println(result)
	}

	qqqq, rrrr:= div(13, 6)
	fmt.Println(qqqq, rrrr)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
