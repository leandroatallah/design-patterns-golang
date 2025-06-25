package main

import "fmt"

type Type int

const (
	Car Type = iota
	Motorcycle
	Bicycle
)

type Vehicle struct {
	Type  Type
	Price float64
}

type TotalTaxCounter interface {
	GetTotalWithTax(t Type, p float64) float64
}

type CarPricer struct{}

func (c CarPricer) GetTotalWithTax(v *Vehicle) float64 {
	return v.Price + (v.Price * 0.2)
}

type MotoPricer struct{}

func (c MotoPricer) GetTotalWithTax(v *Vehicle) float64 {
	return v.Price + (v.Price * 0.15)
}

type BicyclePricer struct{}

func (c BicyclePricer) GetTotalWithTax(v *Vehicle) float64 {
	return v.Price + (v.Price * 0.1)
}

func main() {
	car := Vehicle{Car, 15_000}
	moto := Vehicle{Motorcycle, 12_000}
	bicycle := Vehicle{Bicycle, 2_000}

	carPricer := CarPricer{}
	carTotal := carPricer.GetTotalWithTax(&car)

	motoPricer := MotoPricer{}
	motoTotal := motoPricer.GetTotalWithTax(&moto)

	bicyclePricer := BicyclePricer{}
	bicycleTotal := bicyclePricer.GetTotalWithTax(&bicycle)

	fmt.Println("Car Total Price: ", carTotal)
	fmt.Println("Motorcycle Total Price: ", motoTotal)
	fmt.Println("Bicycle Total Price: ", bicycleTotal)
}
