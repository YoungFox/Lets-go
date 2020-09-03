import (
	"bufio"
	"os"
)

func re() {
	file, err := os.Open("./xxx.txt")
	if err != nil {
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
}