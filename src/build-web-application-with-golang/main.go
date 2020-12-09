package main

import (
	"build-web-application-with-golang/add"
	"build-web-application-with-golang/panictest"
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("Hello world!\n")
	var c complex64 = 5 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v\n", c)

	s := "hello"
	d := []byte(s) // 将字符串 s 转换为 []byte 类型
	d[0] = 'c'
	s2 := string(d) // 再转换回 string 类型
	fmt.Printf("%s\n", s2)

	s0 := "hello"
	s0 = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s0)

	m := `hello
					world`

	fmt.Print(m)

	err := errors.New("\nemit macho dwarf: elf header corrupted\n")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("1111111111111111111111")
	go func() {
		panictest.PanicTest()
	}()
	fmt.Print("2222222222222222222222")

	x := 3

	fmt.Println("x = ", x) // 应该输出 "x = 3"

	x1 := add.Add1(&x) //调用add1(x)

	fmt.Println("x+1 = ", *x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)     // 应该输出"x = 3"

}
