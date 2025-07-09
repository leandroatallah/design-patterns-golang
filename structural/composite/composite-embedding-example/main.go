package main

import "fmt"

// Athlete is a basic struct with a Train method.
// This will be used to demonstrate composition and embedding.
type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

// CompositeSwimmerA demonstrates the COMPOSITE pattern via explicit composition.
// It contains an Athlete (as a field) and a Swim function.
// This is not Go embedding, but rather explicit composition.
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming")
}

type Animal struct{}

func (r *Animal) Eat() {
	fmt.Println("Eating")
}

// Shark demonstrates Go EMBEDDING.
// Animal is embedded, so Shark gets Eat() "for free".
// Swim is a composed function field.
type Shark struct {
	Animal
	Swim func()
}

// Swimmer is an interface for types that can swim.
type Swimmer interface {
	Swim()
}

// Trainer is an interface for types that can train.
type Trainer interface {
	Train()
}

// SwimmerImpl is a concrete implementation of Swimmer.
type SwimmerImpl struct{}

// Swim implements the Swimmer interface.
func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming")
}

// CompositeSwimmerB demonstrates the COMPOSITE pattern via interface composition and embedding.
// It embeds Trainer and Swimmer interfaces, so any type that implements these can be used.
type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {
	// Using explicit composition: CompositeSwimmerA
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	// Using embedding: Shark
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()

	// Using interface embedding: CompositeSwimmerB
	swimmerB := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmerB.Train()
	swimmerB.Swim()
}
