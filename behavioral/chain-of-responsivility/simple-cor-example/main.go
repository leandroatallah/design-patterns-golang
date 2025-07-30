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
