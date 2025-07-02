package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)
	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	bob := Person{"Bob", &Address{"123 Street", "Fundão", "Portugal"}, []string{"Luke", "Matthew"}}
	clark := bob.DeepCopy()
	clark.Name = "Clark"
	clark.Address.StreetAddress = "321 Another Street"
	clark.Friends = append(clark.Friends, "Holy")

	fmt.Println(bob, bob.Address)
	fmt.Println(clark, clark.Address)
}
