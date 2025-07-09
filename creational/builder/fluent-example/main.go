package main

import "fmt"

type Person struct {
	City, Country         string
	Position, CompanyName string
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}
func (b *PersonAddressBuilder) At(country string) *PersonAddressBuilder {
	b.person.Country = country
	return b
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonJobBuilder) As(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}
func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func main() {
	personBuilder := NewPersonBuilder()
	personBuilder.
		Lives().In("Rio de Janeiro").At("Brazil").
		Works().As("Programmer").At("TechTechTech")
	result := personBuilder.Build()
	fmt.Println(result)
}
