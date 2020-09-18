package main

import "fmt"

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value: ", count)
}

func add5Pointer(count *int) {
	*count += 5
	fmt.Println("add5Pointer: ", *count)
}

func main() {
	var number int

	add5Value(number)
	fmt.Println("post add5Value: ", number)

	add5Pointer(&number)
	fmt.Println("post add5Pointer: ", number)

	a, b := 5, 10
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}
