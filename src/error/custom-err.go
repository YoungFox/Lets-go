package main

import (
	"fmt"
	"math"
)

type customError struct {
	Num     float64
	problem string
}

func (e customError) Error() string {
	return fmt.Sprintf("%s!!!,because %f is a negative number", e.problem, e.Num)
}

// Sqrt :使用自定义错误
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, customError{Num: f, problem: "不能是负数"}
	}
	return math.Sqrt(f), nil
}

func main() {
	result, err := Sqrt(-10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
