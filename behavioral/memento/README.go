# Memento

> Memento is a behavioral design pattern that lets you save and restore the previous state of an object without revealing the details of its implementation. (Refactoring Guru)

> The memento pattern is a software design pattern that exposes the private internal state of an object. One example of how this can be used is to restore an object to its previous state (undo via rollback), another is versioning, another is custom serialization. (Wikipedia)

## Example

### Basic

```go
package main

import "fmt"

// The State type is the core business object
type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return memento{}, fmt.Errorf("Index not found\n")
	}
	return c.mementoList[i], nil
}

func main() {
	state := State{"Short description"}
	originator := originator{state}

	fmt.Println(state)

	careTaker := careTaker{}
	careTaker.Add(originator.NewMemento())
	state.Description = "New short description"

	fmt.Println(state)

	m, _ := careTaker.Memento(0)
	originator.ExtractAndStoreState(m)
	state = originator.state

	fmt.Println(state)
}
```

### Undo Redo

```go
package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memento
	current int
}

func NewBankAccount(balance int) *BankAccount {
	b := &BankAccount{balance: balance}
	b.changes = append(b.changes, &Memento{balance})
	return b
}

func (b *BankAccount) String() string {
	return fmt.Sprint("Balance = $", b.balance)
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	m := Memento{b.balance}
	b.changes = append(b.changes, &m)
	b.current++
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
	return &m
}

func (b *BankAccount) Restore(m *Memento) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount) Undo() *Memento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func main() {
	ba := NewBankAccount(100)
	fmt.Println(ba)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1", ba)
	ba.Undo()
	fmt.Println("Undo 2", ba)
	ba.Redo()
	fmt.Println("Redo 1", ba)
}
```
