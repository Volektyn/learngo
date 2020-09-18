package main

import "fmt"

func main() {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	newWeek := append(week[len(week)-1:], week[:len(week)-1]...)
	fmt.Println(week)
	fmt.Println(newWeek)

	week[4] = "Changes"
	fmt.Println(week)
	fmt.Println(newWeek)
}
