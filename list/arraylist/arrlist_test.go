package arraylist_test

import (
	"sync"
	"testing"

	"github.com/454270186/GoDazzle/list/arraylist"
)

func TestTFArrList_Concurrency(t *testing.T) {
	// 创建并发测试的 TFArrList 实例
	list := arraylist.New(arraylist.Option{ThreadSafe: true})

	// 设置并发测试的线程数
	numThreads := 1

	// 设置每个线程要执行的操作次数
	operationsPerThread := 5

	// 设置一个等待组，以便在所有线程完成测试前等待
	var wg sync.WaitGroup
	wg.Add(numThreads)

	// 启动并发测试
	for i := 0; i < numThreads; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < operationsPerThread; j++ {
				// 在这里执行对 TFArrList 的操作，例如 Add、Remove、Get、Contains
				// ...

				// 例如，在每个线程中执行 Add 操作
				list.Add(6)
			}
		}()
	}

	// 等待所有线程完成测试
	wg.Wait()

	// 在测试完成后，检查 TFArrList 的状态或进行其他断言
	// ...

	// 例如，检查 TFArrList 的大小是否与预期一致
	expectedSize := int32(numThreads * operationsPerThread)
	actualSize := int32(list.Size())
	if expectedSize != actualSize {
		t.Errorf("Expected size: %d, Actual size: %d\n", expectedSize, actualSize)
		t.Errorf("%v", list.String())
		t.Error(list.Values()...)
	}
}