package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subscriptions *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subscriptions.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subscriptions.Front(); z != nil; z = z.Next() {
		// if z.Value == x {
		if z.Value.(Observer) == x {
			o.subscriptions.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subscriptions.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

// Concrete observer
type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s\n", data.(string))
}

// Concrete observer
type Mother struct{}

func (m *Mother) Notify(data interface{}) {
	fmt.Printf("Mommy says: \"poor boy... let me take care of you %s\"\n", data.(string))
}

func main() {
	p1 := NewPerson("Bruce")
	ds := &DoctorService{}
	mm := &Mother{}
	p1.Subscribe(ds)
	p1.Subscribe(mm)
	p1.CatchACold()
}
