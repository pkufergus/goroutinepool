package goroutinepool

import (
	"time"
	"fmt"
	"testing"
	"sync"
)

func TestNewRoutinePool(t *testing.T) {
	p := NewRoutinePool(10)
	for i := 0; i < 100; i++ {
		count := i
		p.AddJob(func() {
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("%d\n", count)
		})

	}
	p.Wait()
	//time.Sleep(2 * time.Second)
}


func TestSum(t *testing.T) {
	p := NewRoutinePool(10)
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
	p.Wait()
	//time.Sleep(2 * time.Second)
	fmt.Printf("sum=%d\n", g_index)
}