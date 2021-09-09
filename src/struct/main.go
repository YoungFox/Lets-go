// package main

// import "fmt"

// type Human struct {
// 	name   string
// 	age    int
// 	weight int
// }

// type Student struct {
// 	Human
// 	speciality string
// }

// func main() {
// 	mark := Student{Human{"Mark", 25, 120}, "computer science"}

// 	fmt.Println("His name is ", mark.name)
// 	fmt.Println("His age is ", mark.age)
// 	fmt.Println("His weight is ", mark.weight)
// 	fmt.Println("Hist speciality is ", mark.speciality)
// 	mark.speciality = "AI"

// 	fmt.Println("Mark changed his speciality")
// 	fmt.Println("His speciality is ", mark.speciality)
// 	fmt.Println("Mark become old")
// 	mark.age = 46
// 	fmt.Println("His age is ", mark.age)

// }
// package main

// import "fmt"

// type Skills []string

// type Human struct {
// 	name   string
// 	age    int
// 	weight int
// }

// type Student struct {
// 	Human      // 匿名字段，struct
// 	Skills     // 匿名字段，自定义的类型string slice
// 	int        // 内置类型作为匿名字段
// 	speciality string
// }

// func main() {
// 	// 初始化学生Jane
// 	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
// 	// 现在我们来访问相应的字段
// 	fmt.Println("Her name is ", jane.name)
// 	fmt.Println("Her age is ", jane.age)
// 	fmt.Println("Her weight is ", jane.weight)
// 	fmt.Println("Her speciality is ", jane.speciality)
// 	// 我们来修改他的skill技能字段
// 	jane.Skills = []string{"anatomy"}
// 	fmt.Println("Her skills are ", jane.Skills)
// 	fmt.Println("She acquired two new ones ")
// 	jane.Skills = append(jane.Skills, "physics", "golang")
// 	fmt.Println("Her skills now are ", jane.Skills)
// 	// 修改匿名内置类型字段
// 	jane.int = 3
// 	fmt.Println("Her preferred number is", jane.int)
// }

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
