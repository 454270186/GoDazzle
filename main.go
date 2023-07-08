package main

import (
	"fmt"

	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/list/linklist"
)

func main() {
	linklist := linklist.New(2, 1, 6, 5, 3)
	fmt.Println(linklist.Values()...)
	linklist.Sort(cmp.IntComparator)
	fmt.Println(linklist.Values()...)
}