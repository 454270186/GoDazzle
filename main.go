package main

import "github.com/454270186/GoDazzle/list/arraylist"

func main() {
	arrList := arraylist.New(1, 2, 3, 4)
	arrList.Println()
	arrList.Remove(2)
	arrList.Println()
}