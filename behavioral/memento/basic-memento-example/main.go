package main

import "fmt"

// The State type is the core business object
type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return memento{}, fmt.Errorf("Index not found\n")
	}
	return c.mementoList[i], nil
}

func main() {
	state := State{"Short description"}
	originator := originator{state}

	fmt.Println(state)

	careTaker := careTaker{}
	careTaker.Add(originator.NewMemento())
	state.Description = "New short description"

	fmt.Println(state)

	m, _ := careTaker.Memento(0)
	originator.ExtractAndStoreState(m)
	state = originator.state

	fmt.Println(state)
}
