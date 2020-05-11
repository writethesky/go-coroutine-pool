package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestFunc(t *testing.T) {

	poolLen := 5
	// 初始化一个容量为poolLen的协程池
	pools := New(poolLen)
	fmt.Printf("初始化一个容量为%d的协程池\n", poolLen)

	// 模拟10次并发, 使用协程池
	wg.Add(10)
	t1 := time.Now().UnixNano() / 1e6
	for i := 0; i < 10; i++ {
		go pools.AddTask(task)
	}
	wg.Wait()

	fmt.Printf("使用协程池技术，并发10次，所需时间为：%dms\n", time.Now().UnixNano()/1e6-t1)

	// 模拟10次并发, 不使用协程池
	wg.Add(10)
	t2 := time.Now().UnixNano() / 1e6
	for i := 0; i < 10; i++ {
		go task()
	}
	wg.Wait()

	fmt.Printf("不使用协程池技术，并发10次，所需时间为：%dms\n", time.Now().UnixNano()/1e6-t2)

}

func task() {
	// 模拟执行1s的耗时任务
	time.Sleep(time.Second)
	wg.Done()
}
