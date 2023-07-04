package main

import (
	"fmt"
	"sync"

	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/list/arraylist"
)

func main() {
	list := arraylist.New(arraylist.Option{ThreadSafe: true})

	list.Add(5, 4, 3, 2, 1)

	// 设置并发任务的数量
	numTasks := 2

	// 创建等待组，以确保所有任务完成后再继续执行
	var wg sync.WaitGroup
	wg.Add(numTasks)

	go func ()  {
		defer wg.Done()
		fmt.Println(list.Values()...)
		list.Sort(cmp.IntComparator)
		fmt.Println(list.Values()...)
		
	}()

	go func ()  {
		defer wg.Done()
		fmt.Println(list.Values()...)
		list.Clear()
		list.Add(434, 343, 333)
		fmt.Println(list.Values()...)
	}()

	// 等待所有任务完成
	wg.Wait()

	fmt.Println(list.String())
}