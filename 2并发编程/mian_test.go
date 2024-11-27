package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

/**1.通道关闭的案例*/
/**2.单独退出通道，传输数据不使用一个channel*/

/**time.After(time.Second * 2) 超过2秒就要发送信号了 2秒后的就收不到了*/
func Test2(t *testing.T) {
	test := make(chan int, 10)
	go func(info chan int) {
		for {
			select {
			case val := <-info:
				t.Logf("data %d\n", val)
			case <-time.After(time.Second * 2):
				t.Logf("Time out\n")
				return
			}
		}
	}(test)
	go func() {
		test <- 1
		time.Sleep(time.Second * 2)
		test <- 2
	}()
	time.Sleep(time.Second * 5)
}

/**2.单独退出通道，传输数据不使用一个channel*/
func Test1(t *testing.T) {
	test := make(chan int, 5)
	exit := make(chan struct{})
	go func(info chan int, exit chan struct{}) {
		for {
			select {
			case val := <-info:
				t.Logf("data %d\n", val)
			case <-exit:
				t.Logf("Task Exit!!\n")
				return
			}
		}
	}(test, exit)
	go func() {
		test <- 1
		time.Sleep(time.Second * 1)
		test <- 2
		close(exit)
	}()
	time.Sleep(time.Second * 5)
}

/**1.通道关闭的案例*/
func Test(t *testing.T) {
	test := make(chan int, 10)
	go func(info chan int) {
		for {
			select {
			case val, ok := <-info: //通道关闭的时候，这里ok是false
				if !ok {
					t.Logf("Channel closed")
					return
				}
				t.Logf("data %d\n", val)
			}
		}
	}(test)
	go func() {
		test <- 1
		time.Sleep(time.Second * 1)
		test <- 2
		close(test)
	}()
	time.Sleep(time.Second * 5)
}

// WithValue()函数 值可以进行传递   键值对模式
func TestContext(t *testing.T) {
	a := context.Background()
	b := context.WithValue(a, "k1", "val1")
	c := context.WithValue(b, "key1", "val1")
	d := context.WithValue(c, "key2", "val2")
	e := context.WithValue(d, "key3", "val3")
	f := context.WithValue(e, "key3", "val4444")
	fmt.Printf(" %s\n", f.Value("key2"))
	fmt.Printf(" %s\n\n", f.Value("key3"))
}

/**Done()确定上下文是否完成*/
/**而取消上下文则是最直接的方式，之前进行了context.WithCancel*/
func TestContext1(t *testing.T) {
	//ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done(): //管道关闭,接收取消信号
				t.Logf("Context cancelled")
				return
			}
		}
	}()
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()
	time.Sleep(time.Second * 20)
}
func TestContext2(t *testing.T) {
	ctxA, cancelA := context.WithCancel(context.Background())

	ctxB, cancelB := context.WithCancel(ctxA)
	ctxC, cancelC := context.WithCancel(ctxA)
	ctxD, _ := context.WithCancel(ctxA)

	ctxE, _ := context.WithCancel(ctxB)
	ctxF, _ := context.WithCancel(ctxB)

	ctxG, _ := context.WithCancel(ctxC)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go task("A", ctxA, &wg)

	wg.Add(1)
	go task("B", ctxB, &wg)

	wg.Add(1)
	go task("C", ctxC, &wg)

	wg.Add(1)
	go task("D", ctxD, &wg)

	wg.Add(1)
	go task("E", ctxE, &wg)

	wg.Add(1)
	go task("F", ctxF, &wg)

	wg.Add(1)
	go task("G", ctxG, &wg)

	time.Sleep(2 * time.Second)

	cancelB()
	time.Sleep(1 * time.Second)

	cancelC()
	time.Sleep(1 * time.Second)

	cancelA()
	time.Sleep(1 * time.Second)

	wg.Wait()
	fmt.Println("All tasks stopped")

}
func task(name string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %s started\n", name)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Task %s stopped\n", name)
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}
