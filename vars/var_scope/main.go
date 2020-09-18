package main

import "fmt"

func main() {
	var message string
	count := 5
	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not greater than 5"
	}
	fmt.Println(message)

	count = 0
	if count < 5 {
		count = 10
		count++
	}
	fmt.Println(count == 11)
}
