package main

// #include "person.ext.h"

//so:extern
type Account struct {
	name    string
	balance int64
	flags   []uint8
}

func account_inc_balance(acc *Account, amount int64) int64

func main() {
	acc := Account{
		name:    "Alice",
		balance: 100,
		flags:   []uint8{42},
	}

	balBefore := account_inc_balance(&acc, 50)

	println(
		"name =", acc.name,
		"balance =", balBefore, acc.balance,
		"flags[0] =", acc.flags[0],
	)
}
