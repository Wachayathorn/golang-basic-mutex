package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int64
)

func init() {
	balance = 100
}

func deposit(value int64) {
	mutex.Lock()
	defer mutex.Unlock()
	balance += value
	fmt.Printf("Depositing %d to account with balance : %d \n", value, balance)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		deposit(200)
		wg.Done()
	}()

	go func() {
		deposit(500)
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("Balance is %d\n", balance)
}
