package main

import (
	"fmt"
	"sort"
)

func main() {
	names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names)
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
