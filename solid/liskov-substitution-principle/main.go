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
