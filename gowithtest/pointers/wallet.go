package main

import (
	"errors"
	"fmt"
)

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ types & variables def ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ errors def ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

/*ஐఴஐ๑ஐఴஐஐஐఴஐ๑ஐఴஐஐஐఴ
಄ะ methods def ะ಄
ஐஐळஐ๑ஐळஐஐஐळஐ๑ஐळஐஐஐळ*/

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC ", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

/*

fmt.Printf("address of balance in Deposit is %v \n", &w.balance)

*/
