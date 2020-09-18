package main

import "fmt"

func main() {
	comment1 := `This is the BEST
thing ever!`
	comment2 := `This is the BEST\nthing ever!`
	comment3 := "This is the BEST\nthing ever!"
	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")

	username := "Sir_King_Über"
	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}
	fmt.Print("\n")
	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]), " ")
	}
	fmt.Print("\n")
	runes := []rune(username)
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
	fmt.Print("\n")

	logLevel := "デバッグ"
	for offset, val := range logLevel {
		fmt.Println(offset, string(val))
	}
	fmt.Print("\n")
}
