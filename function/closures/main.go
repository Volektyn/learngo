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

// 1.Create a function called nonLoggedHours() func(int) int. Each time this function is
// called, it will calculate the hours of the employee that have not been logged. You
// will be using a closure inside the function.
func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += 1
		return total
	}
}

// 2. Create a method called PayDay()(int,bool). This method will calculate the weekly
// pay. It needs to take into consideration overtime pay. The method will pay twice
// the hourly rate for hours greater than 40. The function will return int as the weekly
// pay and bool for if the pay is overtime pay. The Boolean will be true if the employee
// worked more than 40 hours and false if they worked less than 40 hours.
func (d *Developer) PayDay() (int, bool) {
	if d.HoursWorked() > 40 {
		overtimeHours := d.HoursWorked() - 40
		overtimePayment := overtimeHours * 2 * d.HourlyRate
		regularPayment := d.HoursWorked() * d.HourlyRate
		return regularPayment + overtimePayment, true
	}
	return d.HoursWorked() * d.HourlyRate, false
}

// 3. Create a method called PayDetails(). This method will print each day and the hours
// worked that day by the employee. It will print the total hours for the week, the pay
// for the week, and if the pay contains overtime pay.
func (d *Developer) PayDetails() {
	for i, v := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours: ", v)
		case 2:
			fmt.Println("Tuesday hours: ", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours: ", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
	}
	fmt.Printf("\nHours worked this week:  %d\n", d.HoursWorked())
	pay, overtime := d.PayDay()
	fmt.Println("Pay for the week: $", pay)
	fmt.Println("Is this overtime pay: ", overtime)
	fmt.Println()
}

// 4. Inside of the main function, initialize a variable of type Developer. Assign a variable
// to nonLoggedHours. Print the variable assigned to nonLoggedHours with values of 2, 3,
// and 5.
// 5.
// Also, in the main() function, log the hours for the following days: Monday 8,
// Tuesday 10, Wednesday 10, Thursday 10, Friday 6, and Saturday 8.
// 6. Then run the PayDetails() method.
func main() {
	dev := Developer{Individual: Employee{ID: 1, FirstName: "Volek", LastName: "1"}, HourlyRate: 15}
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5), "\n")
	dev.LogHours(Monday, 8)
	dev.LogHours(Tuesday, 10)
	dev.LogHours(Wednesday, 10)
	dev.LogHours(Thursday, 10)
	dev.LogHours(Friday, 6)
	dev.LogHours(Saturday, 8)
	dev.PayDetails()
}
