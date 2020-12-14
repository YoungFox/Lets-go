// package main

// import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	var P person // P现在就是person类型的变量了

// 	P.name = "Astaxie"                            // 赋值"Astaxie"给P的name属性.
// 	P.age = 25                                    // 赋值"25"给变量P的age属性
// 	fmt.Printf("The person's name is %s", P.name) // 访问P的name属性.
// }

// package main

// import "fmt"

// // 声明一个新的类型
// type person struct {
// 	name string
// 	age  int
// }

// // 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// // struct也是传值的
// func Older(p1, p2 person) (person, int) {
// 	if p1.age > p2.age { // 比较p1和p2这两个人的年龄
// 		return p1, p1.age - p2.age
// 	}
// 	return p2, p2.age - p1.age
// }

// func main() {
// 	var tom person

// 	// 赋值初始化
// 	tom.name, tom.age = "Tom", 18

// 	// 两个字段都写清楚的初始化
// 	bob := person{age: 25, name: "Bob"}

// 	// 按照struct定义顺序初始化值
// 	paul := person{"Paul", 43}

// 	tb_Older, tb_diff := Older(tom, bob)
// 	tp_Older, tp_diff := Older(tom, paul)
// 	bp_Older, bp_diff := Older(bob, paul)

// 	fmt.Printf("Of %s and %s, %s is older by %d years\n",
// 		tom.name, bob.name, tb_Older.name, tb_diff)

// 	fmt.Printf("Of %s and %s, %s is older by %d years\n",
// 		tom.name, paul.name, tp_Older.name, tp_diff)

// 	fmt.Printf("Of %s and %s, %s is older by %d years\n",
// 		bob.name, paul.name, bp_Older.name, bp_diff)
// }

// package main

// import "fmt"

// type Human struct {
// 	name   string
// 	age    int
// 	weight int
// }

// type Student struct {
// 	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
// 	speciality string
// }

// func main() {
// 	// 我们初始化一个学生
// 	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

// 	// 我们访问相应的字段
// 	fmt.Println("His name is ", mark.name)
// 	fmt.Println("His age is ", mark.age)
// 	fmt.Println("His weight is ", mark.weight)
// 	fmt.Println("His speciality is ", mark.speciality)
// 	// 修改对应的备注信息
// 	mark.speciality = "AI"
// 	fmt.Println("Mark changed his speciality")
// 	fmt.Println("His speciality is ", mark.speciality)
// 	// 修改他的年龄信息
// 	fmt.Println("Mark become old")
// 	mark.age = 46
// 	fmt.Println("His age is", mark.age)
// 	// 修改他的体重信息
// 	fmt.Println("Mark is not an athlet anymore")
// 	mark.weight += 60
// 	fmt.Println("His weight is", mark.weight)
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

// package main

// import "fmt"

// type Human struct {
// 	name  string
// 	age   int
// 	phone string // Human类型拥有的字段
// }

// type Employee struct {
// 	Human      // 匿名字段Human
// 	speciality string
// 	phone      string // 雇员的phone字段
// }

// func main() {
// 	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
// 	fmt.Println("Bob's work phone is:", Bob.phone)
// 	// 如果我们要访问Human的phone字段
// 	fmt.Println("Bob's personal phone is:", Bob.Human.phone)

// 	fmt.Println(Bob.age)
// 	fmt.Println(Bob.Human.age)
// 	fmt.Println(Bob.age == Bob.Human.age)
// }

// package main

// import "fmt"

// const (
// 	WHITE = iota
// 	BLACK
// 	BLUE
// 	RED
// 	YELLOW
// )

// type Color byte

// type Box struct {
// 	width, height, depth float64
// 	color                Color
// }

// type BoxList []Box //a slice of boxes

// func (b Box) Volume() float64 {
// 	return b.width * b.height * b.depth
// }

// func (b *Box) SetColor(c Color) {
// 	b.color = c
// }

// func (bl BoxList) BiggestColor() Color {
// 	v := 0.00
// 	k := Color(WHITE)
// 	for _, b := range bl {
// 		if bv := b.Volume(); bv > v {
// 			v = bv
// 			k = b.color
// 		}
// 	}
// 	return k
// }

// func (bl BoxList) PaintItBlack() {
// 	for i := range bl {
// 		bl[i].SetColor(BLACK)
// 	}
// }

// func (c Color) String() string {
// 	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
// 	return strings[c]
// }

// func main() {
// 	fmt.Println(YELLOW)
// 	boxes := BoxList{
// 		Box{4, 4, 4, RED},
// 		Box{10, 10, 1, YELLOW},
// 		Box{1, 1, 20, BLACK},
// 		Box{10, 10, 1, BLUE},
// 		Box{10, 30, 1, WHITE},
// 		Box{20, 20, 20, YELLOW},
// 	}

// 	fmt.Printf("We have %d boxes in our set\n", len(boxes))
// 	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
// 	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
// 	fmt.Println("The biggest one is", boxes.BiggestColor().String())

// 	fmt.Println("Let's paint them all black")
// 	boxes.PaintItBlack()
// 	fmt.Println("The color of the second one is", boxes[1].color.String())

// 	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
// }

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
