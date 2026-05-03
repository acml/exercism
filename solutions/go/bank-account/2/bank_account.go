package account

import "sync"

// Account represents a bank account
type Account struct {
	balance  int64
	isClosed bool
	m        *sync.Mutex
}

// Open creates a new bank account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		balance: initialDeposit,
		m:       &sync.Mutex{},
	}
}

// Close the bank account
func (a *Account) Close() (payout int64, ok bool) {
	a.m.Lock()
	defer a.m.Unlock()
	if a.isClosed {
		return 0, false
	}
	a.isClosed = true
	return a.balance, true
}

// Balance returns back the balance of the bank account
func (a *Account) Balance() (balance int64, ok bool) {
	if a.isClosed {
		return 0, false
	}
	return a.balance, true
}

// Deposit is used for both bank account deposit and withdrawals.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if a.isClosed {
		return 0, false
	}
	a.m.Lock()
	defer a.m.Unlock()
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
