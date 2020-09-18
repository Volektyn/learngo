package payroll

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear int
	Review            map[string]interface{}
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (dev Developer) Pay() (string, float64) {
	return dev.Individual.FirstName + " " + dev.Individual.LastName, dev.HourlyRate * float64(dev.HoursWorkedInYear)
}

func (man Manager) Pay() (string, float64) {
	return man.Individual.FirstName + " " + man.Individual.LastName, man.Salary * man.CommissionRate
}

type Payer interface {
	Pay() (string, float64)
}

func PayDetails(p Payer) {
	fullName, payment := p.Pay()
	fmt.Println(fullName)
	fmt.Println(payment)
}

func (dev *Developer) CalcReviewRating() (float64, error) {
	total := 0
	for _, v := range dev.Review {
		rating, err := getRating(v)
		if err != nil {
			return 0, err
		}

		total += rating
	}
	return float64(total) / float64(len(dev.Review)), nil
}

func getRating(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case string:
		return getStringRating(v)
	default:
		return 0, errors.New("unknown type")
	}
}

func getStringRating(value string) (int, error) {
	switch strings.ToLower(value) {
	case ("excellent"):
		return 5, nil
	case ("good"):
		return 4, nil
	case ("fair"):
		return 3, nil
	case ("poor"):
		return 2, nil
	case ("unsatisfactory"):
		return 1, nil
	default:
		return 0, errors.New("invalid rate: " + value)
	}
}
