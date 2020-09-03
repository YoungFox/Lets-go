package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./xxx.txt")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}

	defer file.Close()

	var buf [128]byte
	var content []byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read file err", err)
		}

		content = append(content, buf[:n]...)
	}

	fmt.Println(string(content))
}
