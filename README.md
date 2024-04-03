# Go Packages for Internship Application

This repository contains two packages, `strings` and `wallet`, which were created as part of internship application tasks.


## strings
The `strings` package provides functions for working with strings in the Go programming language.

### Functions
1. `Reverse`: Reverses a string.
2. `SymbolCount`: Returns the number of symbols in a string.



## wallet
The `wallet` package represents an implementation of a wallet for handling bitcoins in the Go programming language.

### Functions
1. `Deposit(amount Bitcoin)`: Deposits the specified amount of bitcoins into the wallet.
2. `Withdraw(amount Bitcoin) error`: Withdraws the specified amount of bitcoins from the wallet. Returns an ErrInsufficientFunds error if there are insufficient funds in the wallet.
3. `Balance() Bitcoin`: Returns the current balance of the wallet in bitcoins.
