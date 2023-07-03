package cmp

/*
	return a int number:
	- if a < b ===> -1 (or other positive number)
	- if a = b ===> 0
	- if a > b ===> 1 (or other negative number)
*/
type comparator func(a, b any) int