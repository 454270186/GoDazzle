package cmp

/*
	return a int number:
	- if a < b ===> -1 (or other positive number)
	- if a = b ===> 0
	- if a > b ===> 1 (or other negative number)
*/
type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
	num1 := a.(int)
	num2 := b.(int)

	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	}

	return 0
}

func Int8Comparator(a, b interface{}) int {
	num1 := a.(int8)
	num2 := b.(int8)

	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	}

	return 0
}

func Int32Comparator(a, b interface{}) int {
	num1 := a.(int32)
	num2 := b.(int32)

	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	}

	return 0
}

func Int64Comparator(a, b interface{}) int {
	num1 := a.(int64)
	num2 := b.(int64)

	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	}

	return 0
}

func Float32Comparator(a, b interface{}) int {
	num1 := a.(float32)
	num2 := b.(float32)

	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	}

	return 0
}

func Float64Comparator(a, b interface{}) int {
	num1 := a.(float64)
	num2 := b.(float64)

	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	}

	return 0
}


