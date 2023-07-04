package main

import (
	"fmt"
	"sync"

	"github.com/454270186/GoDazzle/list/arraylist"
)

func main() {
	arr := arraylist.New(arraylist.Option{ThreadSafe: true})
	
	var wg sync.WaitGroup
	threadNum := 10
	optNum := 5
	wg.Add(threadNum)
	mulAdd := func ()  {
		defer wg.Done()
		for i := 0; i < optNum; i++ {
			arr.Add(6)
		}
	}

	for i := 0; i < threadNum; i++ {
		go mulAdd()
	}

	wg.Wait()

	fmt.Println(arr.String())
	fmt.Println(arr.Size())
}