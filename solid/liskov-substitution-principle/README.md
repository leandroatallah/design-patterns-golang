# LSP - Liskov Substitution Principle

> An object (such as a class) may be replaced by a sub-object (such as a class that extends the first class) without breaking the program.

- Derived classes need to keep the base class behaviors
- A bad example: A derived class that changes or invalidated functionalities of the base class.

### ❌ Wrong use

```go
package main

import "fmt"

type Order struct {
 items   []string
 isPayed bool
}

func (o *Order) PrintItems() string {
 return strings.Join(o.items, "\n")
}

func (o *Order) MarkAsPayed() {
 o.isPayed = true
}

type DraftOrder struct {
 Order
}

// It makes DraftOrder unable to replace Order and keep the same behavior
func (o *DraftOrder) MarkAsPayed() {
 panic("Draft order can't be marked as payed.")
}

func main() {
 order := Order{[]string{"Pizza", "Bottle of wine"}, false}
 fmt.Println("Is order payd?", order.isPayed)
 order.MarkAsPayed()
 fmt.Println("Is order payd?", order.isPayed)

 draft := DraftOrder{Order{[]string{"Hotdog", "Coca-Cola"}, false}}
 fmt.Println("Is order payd?", draft.isPayed)
 draft.MarkAsPayed()
 fmt.Println("Is order payd?", draft.isPayed)
}
```

### ✅ Good use

```go
package main

import (
 "fmt"
 "strings"
)

// Make order a base class without isPayed funcionality
type Order struct {
 items []string
}

func (o *Order) PrintItems() string {
 return strings.Join(o.items, "\n")
}

// Create ConfirmedOrder and DraftOrder derived from Order
type ConfirmedOrder struct {
 Order
 isPayed bool
}

func (o *ConfirmedOrder) MarkAsPayed() {
 o.isPayed = true
}

type DraftOrder struct {
 Order
}

func main() {
 order := ConfirmedOrder{Order{[]string{"Pizza", "Bottle of wine"}}, false}
 order.MarkAsPayed()
 fmt.Println("Order 1")
 fmt.Println(order.PrintItems())

 fmt.Println()
 fmt.Println("Order 2")
 draft := DraftOrder{Order{[]string{"Hotdog", "Coca-Cola"}}}
 fmt.Println(draft.PrintItems())
}

```
