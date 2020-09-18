package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
	fmt.Println("a: ", *a, " b: ", *b)
}
