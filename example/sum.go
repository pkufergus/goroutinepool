package main

import (
	"github.com/pkufergus/goroutinepool"
	"time"
	"fmt"
	"sync"
)

func main() {

	p := goroutinepool.NewRoutinePool(10)
	g_index := 0
	var lock *sync.Mutex
	lock = new(sync.Mutex)
	for i := 0; i < 100; i++ {
		count := i
		p.AddJob(func() {
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("%d\n", count)
			lock.Lock()
			defer lock.Unlock()
			g_index++
		})

	}
	p.WaitAll()
	fmt.Printf("sum=%d\n", g_index)
}