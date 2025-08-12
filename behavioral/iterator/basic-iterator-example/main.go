package main

import "fmt"

// Model
type Person interface {
	Name() string
}
type Employee struct {
	name string
}

func (e *Employee) Name() string {
	return e.name
}

// iterator interface
type Iterator interface {
	GetNext() (Person, bool)
	HasMore() bool
}

// Concrete iterator
type EmployeeIterator struct {
	collection      *EmployeeCollection
	currentPosition int
}

func NewPersonIterator(collection *EmployeeCollection) *EmployeeIterator {
	return &EmployeeIterator{collection: collection}
}

func (i *EmployeeIterator) GetNext() (Person, bool) { // change return type
	if i.HasMore() {
		result := i.collection.employeeList[i.currentPosition]
		i.currentPosition++
		return result, true
	}
	return nil, false
}
func (i *EmployeeIterator) HasMore() bool {
	return i.currentPosition < len(i.collection.employeeList)
}

// iterable collection
type IterableCollection interface {
	CreateIterator() Iterator
}

// concrete collection
type EmployeeCollection struct {
	employeeList []Person
}

func NewEmployeeCollection(employees []Person) *EmployeeCollection {
	return &EmployeeCollection{employeeList: employees}
}

func (c *EmployeeCollection) CreateIterator() Iterator {
	return NewPersonIterator(c)
}

func main() {
	employees := []Person{
		&Employee{"Leonard"},
		&Employee{"Raphael"},
	}
	iterator := NewEmployeeCollection(employees).CreateIterator()

	for iterator.HasMore() {
		if next, exists := iterator.GetNext(); exists {
			fmt.Println(next.Name())
		}
	}
}
