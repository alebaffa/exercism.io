package account

import "sync"

type Account struct {
	balance int64
	close   bool
	sync.RWMutex
}

func (account *Account) Balance() (balance int64, ok bool) {
	account.RLock()
	balance, ok = account.balance, true
	account.RUnlock()
	return
}

func (account *Account) Close() (payout int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if account.close {
		return 0, false
	}
	account.close = true
	if account.balance < 0 {
		return 0, true
	}
	payout = account.balance
	account.balance = 0

	return payout, true
}

func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	return 0, true
}

func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}
	return &Account{balance: initalDeposit, close: false}
}
