package bank_test

import (
	"testing"

	bank "github.com/DavidMJChoi/goexp/src/goplch9/bank1"
)

func TestBasics(t *testing.T) {
	bank.Deposit(300)
	bank.Withdraw(150)
	balance := bank.Balance()
	expected := 150
	if balance != expected {
		t.Errorf("Basic testgin failed. Expected %d, Got %d", expected, balance)
	}
}
