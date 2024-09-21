package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func main() {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Final Balance: ", account.GetBalance())
}

//saat ini sleep 3 s sudah aman sampai 10rb.
//tp gimana kalo misal goroutine nya kita nggatau brpa banyaknya,
//kita bisa pake waitgroup. next
