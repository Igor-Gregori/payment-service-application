package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	cvv             int
}

func NewCreditCard(number string, name string, expirationMonth int, expirationYear int, cvv int) (*CreditCard, error) {
	cc := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             cvv,
	}
	err := cc.IsValid()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *CreditCard) IsValid() error {
	errNumber := c.validateNumber()
	if errNumber != nil {
		return errNumber
	}
	errMonth := c.validateMonth()
	if errMonth != nil {
		return errMonth
	}
	errYear := c.validateYear()
	if errYear != nil {
		return errYear
	}
	return nil
}

func (c *CreditCard) validateNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !re.MatchString(c.number) {
		return errors.New("invalid credit card number")
	}
	return nil
}

func (c *CreditCard) validateMonth() error {
	if c.expirationMonth > 0 && c.expirationMonth < 13 {
		return nil
	}
	return errors.New("invalid expiration month")
}

func (c *CreditCard) validateYear() error {
	if c.expirationYear >= time.Now().Year() {
		return nil
	}
	return errors.New("invalid expiration year")
}
