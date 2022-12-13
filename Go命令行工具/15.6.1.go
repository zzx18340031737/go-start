package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

const Num int = 10000

func main() {
	f, _ := os.Create("cpu_file.prof")

	//开始CPU Profile ，结果写到文件f中
	pprof.StartCPUProfile(f)
	//结束profile
	defer pprof.StopCPUProfile()

	var str string
	for i := 0; i < Num; i++ {
		str = fmt.Sprintf("%s%d", str, i)
	}
}
