package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	mu    sync.Mutex
	money int
}

func (a *Account) fillDefault() {
	a.money = 500
}

func (a *Account) updateBalance(val int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.money += val
}

func (a *Account) depositMoney(val int) {
	a.updateBalance(val)
}

func (a *Account) withdrawMoney(val int) {
	if a.money < val {
		fmt.Printf("Withdrawal of amount %v failed due to insufficient balance\n", val)
		return
	} else {
		a.updateBalance(-1 * val)
	}
}

func main() {
	createAccount := func() Account {
		var a Account
		a.fillDefault()
		return a
	}
	tempAcc := createAccount()
	deposits := []int{34, 54, 23, 45, 56, 43, 23, 45, 6}
	withdrawls := []int{45, 43, 65, 43, 541, 34, 23, 45, 65}
	updateDeposits := func(a []int) {
		for i := range a {
			tempAcc.depositMoney(deposits[i])
		}
	}
	updateWithdrawals := func(a []int) {
		for i := range a {
			tempAcc.withdrawMoney(withdrawls[i])
		}
	}
	go updateDeposits(deposits)
	go updateWithdrawals(withdrawls)
	time.Sleep(time.Second)
	fmt.Println(tempAcc.money)

}
