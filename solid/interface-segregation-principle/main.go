package main

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type ModernPrinter struct{}

func (m *ModernPrinter) Print() {
	fmt.Println("Printing...")
}

func (m *ModernPrinter) Scan() {
	fmt.Println("Scanning...")
}

type OldPrinter struct{}

func (o *OldPrinter) Print() {
	fmt.Println("Printing...")
}

func main() {
	m := ModernPrinter{}
	m.Print()
	m.Scan()

	o := OldPrinter{}
	o.Print()
}
