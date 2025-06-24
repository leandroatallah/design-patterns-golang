package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Stock struct {
	products []string
}

func (s *Stock) AddProduct(text string) {
	s.products = append(s.products, text)
}

func (s *Stock) String() string {
	return strings.Join(s.products, "\n")
}

// ✅ Using a separated module for persistence
type StockPersistence struct {
	lineSeparator string
}

func (p *StockPersistence) SaveToFile(s *Stock, filename string) error {
	err := os.WriteFile(filename, []byte(
		strings.Join(s.products, p.lineSeparator),
	), 0644)
	return err
}

func main() {
	s := Stock{products: []string{}}
	s.AddProduct("Tomato")
	s.AddProduct("Orange")

	p := StockPersistence{lineSeparator: "\n"}
	err := p.SaveToFile(&s, "products.txt")
	if err != nil {
		err = errors.New("Error on create the file.")
		fmt.Println(err)
		return
	}

	fmt.Println("Everything works fine.")
}
