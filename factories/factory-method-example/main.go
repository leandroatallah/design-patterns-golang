package main

import "fmt"

/*
* Employee (Product) interface: Define the common interface for the objects the factory method creates.
* EmployeeCreator (Abstract Creator): Declares the factory method, which concreate creators will implement.
* Concrete Creators: Implement the EmployeeCreator interface and define which concrete Employee type to create.
* Concrete product: Implement the Employee interface, representing specific types of employee.
 */

// Abstract Creator
type EmployeeCreator interface {
	CreateEmployee(name string) Employee
}

// Concreate Creator
type DeveloperCreator struct {
}

func (c *DeveloperCreator) CreateEmployee(name string) Employee {
	return &DeveloperEmployee{name}
}

func NewDeveloperCreator() *DeveloperCreator {
	return &DeveloperCreator{}
}

type ManagerCreator struct {
	name string
}

func (c *ManagerCreator) CreateEmployee(name string) Employee {
	return &ManagerEmployee{name}
}
func NewManagerCreator() *ManagerCreator {
	return &ManagerCreator{}
}

// Product interface
type Employee interface {
	GetName() string
	SetName(name string)
}

// Product concrete
type DeveloperEmployee struct {
	name string
}

func (e *DeveloperEmployee) GetName() string {
	return fmt.Sprintf("Developer name: %s", e.name)
}
func (e *DeveloperEmployee) SetName(name string) {
	e.name = name
}

type ManagerEmployee struct {
	name string
}

func (e *ManagerEmployee) GetName() string {
	return fmt.Sprintf("Manager name: %s", e.name)
}
func (e *ManagerEmployee) SetName(name string) {
	e.name = name
}

func main() {
	var devCreator EmployeeCreator = NewDeveloperCreator()
	developer := devCreator.CreateEmployee("William")
	fmt.Println(developer.GetName())

	var managerCreator EmployeeCreator = NewManagerCreator()
	manager := managerCreator.CreateEmployee("Edward")
	fmt.Println(manager.GetName())
}
