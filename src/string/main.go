// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	fmt.Println(strings.Contains("seafood", "foo"))
// 	fmt.Println(strings.Contains("seafood", "bar"))
// 	fmt.Println(strings.Contains("seafood", ""))
// 	fmt.Println(strings.Contains("", ""))

// 	s := []string{"foo", "bar", "baz"}
// 	fmt.Println(strings.Join(s, ", "))
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	str := make([]byte, 0, 100)
// 	str = strconv.AppendInt(str, 4567, 10)
// 	str = strconv.AppendBool(str, false)
// 	str = strconv.AppendQuote(str, "abcdefg")
// 	str = strconv.AppendQuoteRune(str, '单')
// 	fmt.Println(string(str))
// }

package main

import (
	"fmt"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
func main() {
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e)
}
