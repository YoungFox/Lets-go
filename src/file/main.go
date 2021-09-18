// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	os.Mkdir("ywl", 0777)
// 	os.MkdirAll("ywl/test1/test2", 0777)
// 	err := os.Remove("ywl")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	os.RemoveAll("ywl")
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	userFile := "ywl.txt"
// 	fout, err := os.Create(userFile)
// 	if err != nil {
// 		fmt.Println(userFile, err)
// 		return
// 	}
// 	defer fout.Close()
// 	for i := 0; i < 10; i++ {
// 		fout.WriteString("Just a test!\r\n")
// 		fout.Write([]byte("Just a test!\r\n"))
// 	}
// }

package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "ywl.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
