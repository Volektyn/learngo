package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}

func main() {
	dev := Developer{Individual: Employee{ID: 1, FirstName: "Volek", LastName: "1"}, HourlyRate: 15}
	dev.LogHours(Monday, 9)
	dev.LogHours(Tuesday, 7)
	fmt.Println(dev.Individual.FirstName, " worked on Monday: ", dev.WorkWeek[Monday], " hours")
	fmt.Println(dev.Individual.FirstName, " worked on Tuesday: ", dev.WorkWeek[Tuesday], " hours")
	fmt.Println(dev.Individual.FirstName, " worked this week: ", dev.HoursWorked(), " hours")
}
