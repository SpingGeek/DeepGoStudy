// @Author Gopher
// @Date 2025/1/29 16:28:00
// @Desc
package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	// 创建 CPU 性能分析文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Error creating CPU profile:", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Error closing CPU profile:", err)
		}
	}(f)

	// 开始 CPU 性能分析
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("Error starting CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1) // 使用 2 个核心
	go a()
	go b()
	time.Sleep(2 * time.Second) // 增加时间，确保能收集到数据

}
