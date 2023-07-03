package main

import (
	"sync"

	"github.com/454270186/GoDazzle/list/arraylist"
)

func main() {
	arr := arraylist.New(arraylist.Option{ThreadSafe: true})
	var wg sync.WaitGroup

	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			arr.Add(1)
			arr.Remove(0)
		}()
	}


	wg.Wait()
}