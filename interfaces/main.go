package main

import (
	"fmt"
	"os"
	"payroll"
)

var employeeReview = make(map[string]interface{})

func init() {
	fmt.Println("Welcome to the Employee Pay and Performance Review")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func init() {
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
}

func main() {

	d := payroll.Developer{Individual: payroll.Employee{ID: 1, FirstName: "Eric", LastName: "Davis"},
		HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := payroll.Manager{Individual: payroll.Employee{ID: 2, FirstName: "Mr.", LastName: "Boss"},
		Salary: 150000, CommissionRate: .07}

	rating, err := d.CalcReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Rating for dev: ", rating)
	payroll.PayDetails(d)
	payroll.PayDetails(m)
}
