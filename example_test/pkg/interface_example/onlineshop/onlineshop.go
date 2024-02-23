package onlineshop

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

func (a *Account) ChangeName(f string, l string) {
	a.FirstName = f
	a.LastName = l
}

func (a Account) String() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}

type Employee struct {
	Account
	Credits float64
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s, Credits: %.2f", e.FirstName, e.LastName, e.Credits)
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits += amount
		return e.Credits, nil
	}
	return 0.0, errors.New("invalid credit amount")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("you can't remove more credits than the account has")
	}
	return 0.0, errors.New("you can't remove negative numbers")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}
