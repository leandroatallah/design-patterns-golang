package main

import "fmt"

var overdraftLimit = -500

// BankAccount is the Receiver in the Command pattern.
// It contains business logic for deposit and withdrawal.
type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)
	}
}

// Command is the Command interface.
// It declares a single method for executing the command.
type Command interface {
	Call()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

// BankAccountCommand is a Concrete Command.
// It stores all information needed to perform an action on a BankAccount.
type BankAccountCommand struct {
	account *BankAccount
	action  Action
	amount  int
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
	case Withdraw:
		b.account.Withdraw(b.amount)
	}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
	cmd2.Call()
	fmt.Println(ba)
}
