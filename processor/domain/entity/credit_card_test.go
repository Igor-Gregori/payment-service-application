package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("any wrong number", "any name", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, errNil := NewCreditCard("5360178235679963", "any name", 12, 2024, 123)
	assert.Nil(t, errNil)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, errMonthExceeded := NewCreditCard("5360178235679963", "any name", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", errMonthExceeded.Error())

	_, errMonth := NewCreditCard("5360178235679963", "any name", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", errMonth.Error())

	_, errNil := NewCreditCard("5360178235679963", "any name", 11, 2024, 123)
	assert.Nil(t, errNil)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, errYear := NewCreditCard("5360178235679963", "any name", 11, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", errYear.Error())
}
