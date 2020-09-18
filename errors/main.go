package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidLastName      = errors.New(" invalid last name ")
	ErrInvalidRoutingNumber = errors.New(" invalid routing number ")
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (deposit *directDeposit) validateRoutingNumber() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if deposit.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
}

func (deposit *directDeposit) validateLastName() error {
	var err error
	validLN := len(strings.TrimSpace(deposit.lastName)) > 0
	if !validLN {
		err = ErrInvalidLastName
	}
	return err
}

func (deposit *directDeposit) report() {
	fmt.Println(deposit.lastName, "\n", deposit.firstName, "\n", deposit.bankName, "\n", deposit.routingNumber, "\n", deposit.accountNumber)
}

func main() {
	dd := directDeposit{
		lastName:      " ",
		firstName:     "Abe",
		bankName:      "XYZ Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}
	dd.validateRoutingNumber()
	err := dd.validateLastName()
	if err != nil {
		fmt.Println(err)
	}
	dd.report()
}
