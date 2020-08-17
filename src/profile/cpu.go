package main

import (
	"time"

	"github.com/pkg/profile"
)

func joinSlice() []string {

	const count = 100000

	var arr []string = make([]string, count)

	for i := 0; i < count; i++ {
		arr[i] = "arr"
	}

	return arr
}
func main() {
	// 开始性能分析, 返回一个停止接口
	stopper := profile.Start(profile.CPUProfile, profile.ProfilePath("."))

	// 在main()结束时停止性能分析
	defer stopper.Stop()

	// 分析的核心逻辑
	joinSlice()

	// 让程序至少运行1秒
	time.Sleep(time.Second)
}
