package panictest

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func PanicTest() {
	fmt.Print(user)
	panic("no value for $USER")
}
