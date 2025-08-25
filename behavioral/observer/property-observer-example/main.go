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

type PropertyChange struct {
	Name  string
	Value interface{}
}

type Person struct {
	Observable
	name string
	age  int
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}
	p.age = age
	p.Fire(PropertyChange{"Age", p.age})
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		name:       name,
		age:        age,
	}
}

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 18 {
			fmt.Println("Congrats, you can drive now!")
			t.o.Unsubscribe(t)
		}
	}
}

func main() {
	p := NewPerson("Bruce", 15)
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)

	for i := 15; i <= 20; i++ {
		fmt.Println("Setting the age to", i)
		p.SetAge(i)
	}
}
