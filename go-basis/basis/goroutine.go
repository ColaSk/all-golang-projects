package basis

import (
	"fmt"
	"sync"
	"time"
)

func Goroutine() {
	// wg用来等待程序完成
	var wg sync.WaitGroup

	testFunc := func(name string, n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			fmt.Println(name, "协程的执行:", i)
			time.Sleep(1 * time.Second)
		}
	}
	wg.Add(2)
	fmt.Println("开始测试")
	fmt.Println("启动协程1")
	go testFunc("func1", 10)
	fmt.Println("启动协程2")
	go testFunc("func2", 10)
	fmt.Println("等待协程执行结束...")
	wg.Wait()

}
