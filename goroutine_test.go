package gokit

import (
	"log/slog"
	"sync"
	"testing"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

// 在 go test 环境中发生panic：
// 1. 不会中断测试执行：当一个 goroutine 发生 panic 而没有被 recover 时，不会导致整个测试程序退出
// 2. 测试会继续运行：其他测试用例会继续执行
// 3. 会记录错误：虽然测试不会中断，但会记录下 panic 信息
//
// 通过设置 GOTRACEBACK=crash 环境变量，可以让go test 在发生 panic 时行为和普通main程序一样
//
//	GOTRACEBACK=crash go test
func TestGo(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	Go(func() {
		defer wg.Done()

		panic("test panic in Go")
	})

	wg.Wait()

	t.Log("TestGo completed")
}

func TestSageGo(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	SafeGo(func() {
		defer wg.Done()

		panic("test panic in SafeGo")
	})

	wg.Wait()

	t.Log("TestSafeGo completed")
}
