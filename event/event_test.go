package event_test

import (
	"context"
	"fmt"
	"github.com/cqu20141693/go-service-common/event"
	"sync"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {

	event.TriggerEvent(event.Start)

	//test timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	timeout := func() {
		hctx, hcancel := context.WithTimeout(ctx, time.Second*4)
		defer hcancel()

		resp := make(chan struct{}, 1)
		// 处理逻辑
		go func() {
			// 处理耗时
			time.Sleep(time.Second * 10)
			resp <- struct{}{}
		}()

		// 超时机制
		select {
		case <-ctx.Done():
			fmt.Println("ctx timeout")
			fmt.Println(ctx.Err())
		case <-hctx.Done():
			fmt.Println("hctx timeout")
			fmt.Println(hctx.Err())
		case v := <-resp:
			fmt.Println("test2 function handle done")
			fmt.Printf("result: %v\n", v)
		}
		fmt.Println("test2 finish")
		return

	}
	go func() {
		event.RegisterHook(event.Start, event.NewHookContext(timeout, "ctx"))
	}()

	time.Sleep(1 * time.Second)

}

func TestCancel(t *testing.T) {
	wg := new(sync.WaitGroup)
	event.TriggerEvent(event.Start)
	ctx, cancel := context.WithCancel(context.Background())
	cancelFunc := func() {
		defer wg.Done()
		respC := make(chan int)
		// 处理逻辑
		go func() {
			time.Sleep(time.Second * 5)
			respC <- 10
		}()
		// 取消机制
		select {
		case <-ctx.Done():
			fmt.Println("cancel")

		case r := <-respC:
			fmt.Println(r)

		}
	}

	wg.Add(1)
	go func() {
		event.RegisterHook(event.Start, event.NewHookContext(cancelFunc, "ctx"))
	}()
	time.Sleep(time.Second * 2)
	// 触发取消
	cancel()
	// 等待goroutine退出
	wg.Wait()
}
