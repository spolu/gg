package account

import ()

type Account struct {
	name string
}

func NewAccount(name string) *Account {
	a := new(Account)
	a.name = name
	return a
}
