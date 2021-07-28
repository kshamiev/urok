// Go рутины, одновременно работающие с общими данными сами собой не могут синхронизироваться
package main

import (
	"fmt"
	"log"
	"sync"
)

var mu sync.RWMutex

// Пусть у нас есть Счет
type Account struct {
	balance float64
}

func (a *Account) Balance() float64 {
	// блокировка на чтение
	// никто не сможет менять данные в исключительной блокировке пока читаем
	// и не можем прочитать пока не изменят данные и не будет снята исключительная блокирвка
	mu.RLock()
	defer mu.RUnlock()
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	// полная блокировка
	// никто не сможет прочитать (также пройти RLock) пока не изменим данные
	// также будем ждать пока все не прочитают и не будет снята блокировка на чтение
	mu.Lock()
	defer mu.Unlock()
	log.Printf("depositing: %f", amount)
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	mu.Lock()
	defer mu.Unlock()
	if amount > a.balance {
		return
	}
	log.Printf("withdrawing: %f", amount)
	a.balance -= amount
}

func main() {
	var wg sync.WaitGroup
	acc := Account{}

	// запускаем отдельные программы
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// Каждая из которых, производит операции с аккаунтом
			for j := 0; j < 100; j++ {
				// Иногда снимает деньги
				if j%2 == 1 {
					acc.Withdraw(50)
					continue
				}
				// иногда кладет
				acc.Deposit(50)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// Что же получится в результате
	fmt.Println(acc.Balance())
}
