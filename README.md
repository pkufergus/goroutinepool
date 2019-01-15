# goroutinepool
Simple Goroutine pool  
一个简单的go golang 协程池  
使用原生的channel和waitgroup实现，简单实用

# How to use

```
go get github.com/pkufergus/goroutinepool
```

## Simple example

```go

package main

import (
	"github.com/pkufergus/goroutinepool"
	"time"
	"fmt"
)

func main() {

	p := goroutinepool.NewRoutinePool(10)
	for i := 0; i < 100; i++ {
		count := i
		p.AddJob(func() {
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("%d\n", count)
		})

	}
	p.WaitAll()
}

```

## Test Parallel Sum

```go

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

```

## Doc


 ## 问题相关

 有任何问题欢迎提Issue,或者将问题描述发送至我邮箱 `caosong1319#gmail.com`.我会尽快解答.推荐提交Issue方式.
