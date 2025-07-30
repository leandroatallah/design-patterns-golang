# Chain of Responsibility

> The chain-of-responsibility pattern is a behavioral design pattern consisting of a source of command objects and a series of processing objects. Each processing object contains logic that defines the types of command objects that it can handle; the rest are passed to the next processing object in the chain. (Wikipedia)

> Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain. (Refactoring Guru)

## Structure

1. The **Handler** declares the interface, common for all concrete handlers. It usually contains just a single method for handling requests, but sometimes it may also have another method for setting the next handler on the chain.
2. The **Base** Handler is an optional class where you can put the boilerplate code that’s common to all handler classes.
3. **Concrete** Handlers contain the actual code for processing requests. Upon receiving a request, each handler must decide whether to process it and, additionally, whether to pass it along the chain.
4. The **Client** may compose chains just once or compose them dynamically, depending on the application’s logic. Note that a request can be sent to any handler in the chain—it doesn’t have to be the first one.

## Examples

```go
package main

import "fmt"

type CustomerBudget struct {
 Approved bool
 Total    int
}

type BaseHandler interface {
 Handle(*CustomerBudget) *CustomerBudget
 SetNextHandler(BaseHandler) BaseHandler
}

// Seller
type SellerBudgetHandler struct {
 nextHandler BaseHandler
}

func (b *SellerBudgetHandler) Handle(budget *CustomerBudget) *CustomerBudget {
 if budget.Total <= 1000 {
  fmt.Println("The seller handled this budget")
  budget.Approved = true
  return budget
 }

 if b.nextHandler != nil {
  return b.nextHandler.Handle(budget)
 }
 return budget
}

func (b *SellerBudgetHandler) SetNextHandler(handler BaseHandler) BaseHandler {
 b.nextHandler = handler
 return handler
}

// Manager
type ManagerBudgetHandler struct {
 nextHandler BaseHandler
}

func (b *ManagerBudgetHandler) Handle(budget *CustomerBudget) *CustomerBudget {
 if budget.Total <= 5000 {
  fmt.Println("The manager handled this budget")
  budget.Approved = true
  return budget
 }

 if b.nextHandler != nil {
  return b.nextHandler.Handle(budget)
 }
 return budget
}

func (b *ManagerBudgetHandler) SetNextHandler(handler BaseHandler) BaseHandler {
 b.nextHandler = handler
 return handler
}

// CEO
type CEOBudgetHandler struct {
 nextHandler BaseHandler
}

func (b *CEOBudgetHandler) Handle(budget *CustomerBudget) *CustomerBudget {
 fmt.Println("The CEO handles any budget")
 budget.Approved = true
 return budget
}

func (b *CEOBudgetHandler) SetNextHandler(handler BaseHandler) BaseHandler {
 b.nextHandler = handler
 return handler
}

func main() {
 // CEO := &CEOBudgetHandler{}
 //
 // manager := &ManagerBudgetHandler{}
 // manager.SetNextHandler(CEO)
 //
 // seller := SellerBudgetHandler{}
 // seller.SetNextHandler(manager)
 //
 // customerBudget := CustomerBudget{Total: 4000}
 // seller.Handle(&customerBudget)

 seller := SellerBudgetHandler{
  nextHandler: &ManagerBudgetHandler{
   nextHandler: &CEOBudgetHandler{},
  },
 }

 customerBudget := CustomerBudget{Total: 2000}
 seller.Handle(&customerBudget)

}

```
