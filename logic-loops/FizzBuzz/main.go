// The rules are as follows:
// • Write a program that prints out the numbers from 1 to 100.
// • If the number is a multiple of 3, print "Fizz."
// • If the number is a multiple of 5, print "Buzz."
// • If the number is a multiple of 3 and 5, print "FizzBuzz."

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	i := 1
	FizzBuzz(i)
}

func FizzBuzz(i int) {
	if i%3 == 0 {
		if i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println("Fizz")
		}
	} else if i%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(strconv.Itoa(i))
	}
	if validateI(i) == nil {
		i++
		FizzBuzz(i)
	}
}

func validateI(i int) error {
	if i < 0 {
		return errors.New("Input can't be a negative number")
	} else if i == 100 {
		return errors.New("Input can't be over 100")
	} else {
		return nil
	}
}
