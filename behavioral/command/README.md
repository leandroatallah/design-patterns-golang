# Command

> Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request. This transformation lets you pass requests as a method arguments, delay or queue a request’s execution, and support undoable operations. (Refactoring Guru)

> The command pattern is a behavioral design pattern in which an object is used to encapsulate all information needed to perform an action or trigger an event at a later time. This information includes the method name, the object that owns the method and values for the method parameters. (Wikipedia)

## Structure

1. The **Sender** class (aka **invoker**) is responsible for initiating requests. This class must have a field for storing a reference to a command object. The sender triggers that command instead of sending the request directly to the receiver. Note that the sender isn’t responsible for creating the command object. Usually, it gets a pre-created command from the client via the constructor.
2. The **Command interface** usually declares just a single method for executing the command.
3. **Concrete Commands** implement various kinds of requests. A concrete command isn’t supposed to perform the work on its own, but rather to pass the call to one of the business logic objects. However, for the sake of simplifying the code, these classes can be merged.
4. The **Receiver** class contains some business logic. Almost any object may act as a receiver. Most commands only handle the details of how a request is passed to the receiver, while the receiver itself does the actual work.
5. The **Client** creates and configures concrete command objects. The client must pass all of the request parameters, including a receiver instance, into the command’s constructor. After that, the resulting command may be associated with one or multiple senders.

## Examples

### Undo command

```go
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

func (b *BankAccount) Withdraw(amount int) bool {
 if b.balance-amount >= overdraftLimit {
  b.balance -= amount
  fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)
  return true
 }
 return false
}

// Command is the Command interface.
// It declares a single method for executing the command.
type Command interface {
 Call()
 Undo()
}

type Action int

const (
 Deposit Action = iota
 Withdraw
)

// BankAccountCommand is a Concrete Command.
// It stores all information needed to perform an action on a BankAccount.
type BankAccountCommand struct {
 account   *BankAccount
 action    Action
 amount    int
 succeeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
 return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (b *BankAccountCommand) Call() {
 switch b.action {
 case Deposit:
  b.account.Deposit(b.amount)
  b.succeeded = true
 case Withdraw:
  b.succeeded = b.account.Withdraw(b.amount)
 }
}

func (b *BankAccountCommand) Undo() {
 if !b.succeeded {
  return
 }
 switch b.action {
 case Deposit:
  b.account.Withdraw(b.amount)
 case Withdraw:
  b.account.Deposit(b.amount)
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
 cmd2.Undo()
 fmt.Println(ba)
}
```

### Composite command

```go
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

func (b *BankAccount) Withdraw(amount int) bool {
 if b.balance-amount >= overdraftLimit {
  b.balance -= amount
  fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)
  return true
 }
 return false
}

// Command is the Command interface.
// It declares a single method for executing the command.
type Command interface {
 Call()
 Undo()
 Succeeded() bool
 SetSucceeded(value bool)
}

type Action int

const (
 Deposit Action = iota
 Withdraw
)

// BankAccountCommand is a Concrete Command.
// It stores all information needed to perform an action on a BankAccount.
type BankAccountCommand struct {
 account   *BankAccount
 action    Action
 amount    int
 succeeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
 return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (b *BankAccountCommand) Call() {
 switch b.action {
 case Deposit:
  b.account.Deposit(b.amount)
  b.succeeded = true
 case Withdraw:
  b.succeeded = b.account.Withdraw(b.amount)
 }
}

func (b *BankAccountCommand) Undo() {
 if !b.succeeded {
  return
 }
 switch b.action {
 case Deposit:
  b.account.Withdraw(b.amount)
 case Withdraw:
  b.account.Deposit(b.amount)
 }
}

func (b *BankAccountCommand) Succeeded() bool {
 return b.succeeded
}
func (b *BankAccountCommand) SetSucceeded(value bool) {
 b.succeeded = value
}

type CompositeBankAccountCommand struct {
 commands []Command
}

func (c *CompositeBankAccountCommand) Call() {
 for _, cmd := range c.commands {
  cmd.Call()
 }

}
func (c *CompositeBankAccountCommand) Undo() {
 for index := range c.commands {
  c.commands[len(c.commands)-index-1].Undo()
 }
}
func (c *CompositeBankAccountCommand) Succeeded() bool {
 for _, cmd := range c.commands {
  if !cmd.Succeeded() {
   return false
  }
 }
 return true
}
func (c *CompositeBankAccountCommand) SetSucceeded(value bool) {
 for _, cmd := range c.commands {
  cmd.SetSucceeded(value)
 }
}

type MoneyTransferCommand struct {
 CompositeBankAccountCommand
 from, to *BankAccount
 amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
 c := &MoneyTransferCommand{from: from, to: to, amount: amount}
 c.commands = append(c.commands, NewBankAccountCommand(from, Withdraw, amount))
 c.commands = append(c.commands, NewBankAccountCommand(to, Deposit, amount))
 return c
}

func (m *MoneyTransferCommand) Call() {
 ok := true
 for _, cmd := range m.commands {
  if ok {
   cmd.Call()
   ok = cmd.Succeeded()
  } else {
   cmd.SetSucceeded(false)
  }
 }
}

func main() {
 from := BankAccount{100}
 to := BankAccount{0}

 mtc := NewMoneyTransferCommand(&from, &to, 25)
 mtc.Call()
 mtc.Undo()

 fmt.Println(from, to)
}
```
