package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

func NewEmployeeFactoryClosure(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct {
	position     string
	annualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.position, f.annualIncome}
}

func NewEmployeeFactory(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	// Closure approach
	employeeFactory := NewEmployeeFactoryClosure("Developer", 80000)
	employee := employeeFactory("Axel")
	fmt.Println(employee)

	// Specialized object approach
	bossFactory := NewEmployeeFactory("CEO", 100000)
	boss := bossFactory.Create("Mark")
	fmt.Println(boss)
}
