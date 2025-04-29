package main

import (
	"fmt"
)

var amountArr = []int{20, 50, 100, 200, 500}

type ATM struct {
	reduce []int
}

func Constructor() ATM {
	return ATM{make([]int, 5)}
}

func (atm *ATM) Deposit(banknotesCount []int) {
	for i := range atm.reduce {
		atm.reduce[i] += banknotesCount[i]
	}
}

func (atm *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		if atm.reduce[i] == 0 {
			continue
		}
		c := amount / amountArr[i]
		c = min(c, atm.reduce[i])
		amount -= c * amountArr[i]
		res[i] = c
	}
	if amount != 0 {
		return []int{-1}
	} else {
		for i := range res {
			atm.reduce[i] -= res[i]
		}
	}
	return res
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */

func main() {
	banknotesCount := []int{0, 0, 1, 2, 1}
	amount := 600
	atm := Constructor()
	atm.Deposit(banknotesCount)
	w := atm.Withdraw(amount)
	fmt.Printf("%v", w)

	atm.Deposit([]int{0, 1, 0, 1, 1})
	w = atm.Withdraw(600)
	fmt.Printf("%v", w)
}
