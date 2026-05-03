package account

import "sync"

// Account represents a bank account
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// Open creates a new bank account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit}
}

// Close the bank account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}

// Balance returns back the balance of the bank account
func (a *Account) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit is used for both bank account deposit and withdrawals.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
