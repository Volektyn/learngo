package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	var (
		maxWord  string
		maxValue int
	)

	for key, value := range words {
		if words[key] > maxValue {
			maxWord = key
			maxValue = value
		}
	}

	fmt.Println("Most popular word: ", maxWord)
	fmt.Println("With a count of: ", maxValue)
}
